<div>

[**简体中文**](README_zh_CN.md)

</div>

## sipclash

[![Downloads](https://img.shields.io/github/downloads/fyy6658-arch/FlClash/total?style=flat-square&logo=github)](https://github.com/fyy6658-arch/FlClash/releases/)[![Last Version](https://img.shields.io/github/release/fyy6658-arch/FlClash/all.svg?style=flat-square)](https://github.com/fyy6658-arch/FlClash/releases/)[![License](https://img.shields.io/github/license/fyy6658-arch/FlClash?style=flat-square)](LICENSE)

[![Upstream Channel](https://img.shields.io/badge/Upstream-Telegram_Channel-blue?style=flat-square&logo=telegram)](https://t.me/FlClash)

A multi-platform proxy client based on ClashMeta, simple and easy to use, open-source and ad-free.

## About This Fork

This is a streamlined personal edition based on
[FlClash](https://github.com/chen08209/FlClash), designed for users who prefer a
focused proxy experience.

It retains the core proxy capabilities and Material You interface of FlClash
while removing non-essential integrations such as Firebase and simplifying
selected features and screens.

### Highlights

- **Streamlined experience**: Reduces non-essential features and third-party
  services while retaining commonly used proxy management capabilities.
- **Automatic proxy selection**: Tests proxy latency for configured selector
  groups and prioritizes low-latency nodes in Taiwan, Singapore, and Japan.
- **Automatic maintenance**: Periodically retests nodes while the app is
  running and attempts to switch when the current node is unavailable or has
  high latency.
- **Open-source and ad-free**: Source code remains available under the GNU GPL
  v3 license.

> [!NOTE]
> This repository is a personal modification of FlClash, not an official
> FlClash release. The fork was first modified in July 2026. See the Git history
> for the complete changes.

on Desktop:
<p style="text-align: center;">
    <img alt="desktop" src="snapshots/desktop.gif">
</p>

on Mobile:
<p style="text-align: center;">
    <img alt="mobile" src="snapshots/mobile.gif">
</p>

## Features

✈️ Multi-platform: Android, Windows, macOS and Linux

💻 Adaptive multiple screen sizes, Multiple color themes available

💡 Based on Material You Design, [Surfboard](https://github.com/getsurfboard/surfboard)-like UI

☁️ Supports data sync via WebDAV

✨ Support subscription link, Dark mode

## Use

### Linux

⚠️ Make sure to install the following dependencies before using them

   ```bash
    sudo apt-get install libayatana-appindicator3-dev
    sudo apt-get install libkeybinder-3.0-dev
   ```

### Android

Support the following actions

   ```bash
    com.fyy6658.sipclash.action.START
    
    com.fyy6658.sipclash.action.STOP
    
    com.fyy6658.sipclash.action.TOGGLE
   ```

## Download

<a href="https://github.com/fyy6658-arch/FlClash/releases"><img alt="Get sipclash on GitHub" src="snapshots/get-it-on-github.svg" width="200px"/></a>

### Homebrew

```bash
brew tap chen08209/tap
brew install --cask flclash
```

## Build

1. Update submodules
   ```bash
   git submodule update --init --recursive
   ```

2. Install `Flutter` and `Golang` environment

3. Build Application

    - android

        1. Install `Android SDK`, `Android NDK`

        2. Set `ANDROID_NDK` environment variable

        3. Run build script

           ```bash
           dart setup.dart android
           ```

    - windows

        1. Requires a Windows client

        2. Install `GCC`, `Inno Setup`

        3. Run build script

           ```bash
           dart setup.dart windows
           ```

    - linux

        1. Requires a Linux client

        2. Dependencies are auto-installed by setup script, or manually:
           ```bash
           sudo apt-get install -y libayatana-appindicator3-dev libkeybinder-3.0-dev
           ```

        3. Run build script

           ```bash
           dart setup.dart linux
           ```

    - macOS

        1. Requires a macOS client

        2. Run build script

           ```bash
           dart setup.dart macos
           ```

## Star

The easiest way to support developers is to click on the star (⭐) at the top of the page.

<p style="text-align: center;">
    <a href="https://api.star-history.com/svg?repos=chen08209/FlClash&Date">
        <img alt="start" width=50% src="https://api.star-history.com/svg?repos=chen08209/FlClash&Date"/>
    </a>
</p>
