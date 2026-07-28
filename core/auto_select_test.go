package main

import "testing"

func TestIsPreferredRegionNode(t *testing.T) {
	matches := []string{
		"🇹🇼 台湾 01",
		"TW-02",
		"新加坡 01",
		"Singapore Premium",
		"🇯🇵 東京 01",
		"JP-03",
	}
	for _, name := range matches {
		if !isPreferredRegionNode(name) {
			t.Fatalf("expected %q to match", name)
		}
	}

	nonMatches := []string{"美国 01", "US-01", "DIRECT", "Japanology"}
	for _, name := range nonMatches {
		if isPreferredRegionNode(name) {
			t.Fatalf("expected %q not to match", name)
		}
	}
}

func TestShouldRefreshAutoSelect(t *testing.T) {
	const testURL = "https://example.com/generate_204"
	tests := []struct {
		name          string
		configuredURL string
		eventURL      string
		delay         uint16
		want          bool
	}{
		{
			name:          "refreshes above threshold",
			configuredURL: testURL,
			eventURL:      testURL,
			delay:         161,
			want:          true,
		},
		{
			name:          "does not refresh at threshold",
			configuredURL: testURL,
			eventURL:      testURL,
			delay:         160,
			want:          false,
		},
		{
			name:          "ignores another test URL",
			configuredURL: testURL,
			eventURL:      "https://example.org/generate_204",
			delay:         300,
			want:          false,
		},
		{
			name:          "refreshes after timeout",
			configuredURL: testURL,
			eventURL:      testURL,
			delay:         0,
			want:          true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := shouldRefreshAutoSelect(
				test.configuredURL,
				test.eventURL,
				test.delay,
			)
			if got != test.want {
				t.Fatalf("shouldRefreshAutoSelect() = %t, want %t", got, test.want)
			}
		})
	}
}
