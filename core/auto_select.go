package main

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/metacubex/mihomo/adapter"
	"github.com/metacubex/mihomo/adapter/outboundgroup"
	"github.com/metacubex/mihomo/common/utils"
	C "github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/tunnel"
)

const (
	autoSelectInterval       = 10 * time.Minute
	autoSelectTimeout        = 8 * time.Second
	autoSelectDelayThreshold = 160
)

var (
	autoSelectLock    sync.Mutex
	autoSelectCancel  context.CancelFunc
	autoSelectContext context.Context
	autoSelectGroups  []string
	autoSelectTestURL string
	autoSelectRunning = map[string]context.Context{}
)

func handleSetAutoSelect(data string) string {
	params := &AutoSelectParams{}
	if err := json.Unmarshal([]byte(data), params); err != nil {
		return err.Error()
	}
	configureAutoSelect(params.AutoSelectGroups, params.TestURL)
	return ""
}

func configureAutoSelect(groupNames []string, testURL string) {
	autoSelectLock.Lock()
	if autoSelectCancel != nil {
		autoSelectCancel()
		autoSelectCancel = nil
	}
	autoSelectContext = nil
	autoSelectGroups = nil
	autoSelectTestURL = ""
	autoSelectRunning = map[string]context.Context{}
	if len(groupNames) == 0 {
		autoSelectLock.Unlock()
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	autoSelectCancel = cancel
	groups := append([]string(nil), groupNames...)
	if testURL == "" {
		testURL = C.DefaultTestURL
	}
	autoSelectContext = ctx
	autoSelectGroups = groups
	autoSelectTestURL = testURL
	autoSelectLock.Unlock()

	go func() {
		runAutoSelect(ctx, groups, testURL)
		ticker := time.NewTicker(autoSelectInterval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				runAutoSelect(ctx, groups, testURL)
			}
		}
	}()
}

func stopAutoSelect() {
	autoSelectLock.Lock()
	defer autoSelectLock.Unlock()
	if autoSelectCancel != nil {
		autoSelectCancel()
		autoSelectCancel = nil
	}
	autoSelectContext = nil
	autoSelectGroups = nil
	autoSelectTestURL = ""
	autoSelectRunning = map[string]context.Context{}
}

type autoSelectResult struct {
	name  string
	delay uint16
}

func runAutoSelect(parent context.Context, groupNames []string, testURL string) {
	for _, groupName := range groupNames {
		runAutoSelectGroup(parent, groupName, testURL)
	}
}

func runAutoSelectGroup(parent context.Context, groupName, testURL string) {
	autoSelectLock.Lock()
	if autoSelectContext != parent {
		autoSelectLock.Unlock()
		return
	}
	if _, running := autoSelectRunning[groupName]; running {
		autoSelectLock.Unlock()
		return
	}
	autoSelectRunning[groupName] = parent
	autoSelectLock.Unlock()

	defer func() {
		autoSelectLock.Lock()
		if autoSelectRunning[groupName] == parent {
			delete(autoSelectRunning, groupName)
		}
		autoSelectLock.Unlock()
	}()
	selectBestRegionalProxy(parent, groupName, testURL)
}

func handleAutoSelectDelay(testURL, proxyName string, delay uint16) {
	autoSelectLock.Lock()
	if !shouldRefreshAutoSelect(autoSelectTestURL, testURL, delay) ||
		autoSelectContext == nil {
		autoSelectLock.Unlock()
		return
	}
	ctx := autoSelectContext
	groupNames := append([]string(nil), autoSelectGroups...)
	configuredTestURL := autoSelectTestURL
	autoSelectLock.Unlock()

	for _, groupName := range groupNames {
		if currentAutoSelectProxy(groupName) == proxyName {
			go runAutoSelectGroup(ctx, groupName, configuredTestURL)
		}
	}
}

func shouldRefreshAutoSelect(configuredTestURL, testURL string, delay uint16) bool {
	return configuredTestURL == testURL &&
		(delay == 0 || delay > autoSelectDelayThreshold)
}

func currentAutoSelectProxy(groupName string) string {
	runLock.Lock()
	rawGroup := tunnel.AllProxies()[groupName]
	runLock.Unlock()
	if rawGroup == nil || rawGroup.Type() != C.Selector {
		return ""
	}
	adapterProxy, ok := rawGroup.(*adapter.Proxy)
	if !ok {
		return ""
	}
	selector, ok := adapterProxy.ProxyAdapter.(interface {
		Now() string
	})
	if !ok {
		return ""
	}
	return selector.Now()
}

func selectBestRegionalProxy(parent context.Context, groupName, testURL string) {
	runLock.Lock()
	rawGroup := tunnel.AllProxies()[groupName]
	runLock.Unlock()
	if rawGroup == nil || rawGroup.Type() != C.Selector {
		return
	}

	adapterProxy, ok := rawGroup.(*adapter.Proxy)
	if !ok {
		return
	}
	group, ok := adapterProxy.ProxyAdapter.(outboundgroup.ProxyGroup)
	if !ok {
		return
	}
	selector, ok := adapterProxy.ProxyAdapter.(outboundgroup.SelectAble)
	if !ok {
		return
	}

	candidates := make([]C.Proxy, 0)
	for _, proxy := range group.Proxies() {
		if isPreferredRegionNode(proxy.Name()) {
			candidates = append(candidates, proxy)
		}
	}
	if len(candidates) == 0 {
		return
	}

	ctx, cancel := context.WithTimeout(parent, autoSelectTimeout)
	defer cancel()
	expectedStatus, err := utils.NewUnsignedRanges[uint16]("")
	if err != nil {
		return
	}

	results := make(chan autoSelectResult, len(candidates))
	var waitGroup sync.WaitGroup
	for _, proxy := range candidates {
		waitGroup.Add(1)
		go func(proxy C.Proxy) {
			defer waitGroup.Done()
			delay, err := proxy.URLTest(ctx, testURL, expectedStatus)
			if err == nil && delay > 0 {
				results <- autoSelectResult{name: proxy.Name(), delay: delay}
			}
		}(proxy)
	}
	waitGroup.Wait()
	close(results)

	best := autoSelectResult{}
	for result := range results {
		if best.delay == 0 || result.delay < best.delay {
			best = result
		}
	}
	if best.delay == 0 {
		return
	}
	if ctx.Err() != nil {
		return
	}
	currentSelector, ok := adapterProxy.ProxyAdapter.(interface {
		Now() string
	})
	if ok && currentSelector.Now() == best.name {
		return
	}
	if err := selector.Set(best.name); err != nil {
		return
	}
	sendMessage(Message{
		Type: AutoSelectMessage,
		Data: map[string]string{
			"group": groupName,
			"proxy": best.name,
		},
	})
}

func isPreferredRegionNode(name string) bool {
	lowerName := strings.ToLower(name)
	for _, keyword := range []string{
		"台湾", "台灣", "🇹🇼",
		"新加坡", "狮城", "獅城", "🇸🇬",
		"日本", "东京", "東京", "大阪", "🇯🇵",
	} {
		if strings.Contains(lowerName, keyword) {
			return true
		}
	}
	tokens := strings.FieldsFunc(lowerName, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for _, token := range tokens {
		switch token {
		case "tw", "taiwan", "sg", "singapore", "jp", "japan":
			return true
		}
	}
	return false
}
