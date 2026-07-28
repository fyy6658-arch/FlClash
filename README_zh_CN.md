<div>

[**English**](README.md)

</div>

## sipclash

[![Downloads](https://img.shields.io/github/downloads/fyy6658-arch/FlClash/total?style=flat-square&logo=github)](https://github.com/fyy6658-arch/FlClash/releases/)[![Last Version](https://img.shields.io/github/release/fyy6658-arch/FlClash/all.svg?style=flat-square)](https://github.com/fyy6658-arch/FlClash/releases/)[![License](https://img.shields.io/github/license/fyy6658-arch/FlClash?style=flat-square)](LICENSE)

[![上游频道](https://img.shields.io/badge/上游-Telegram_频道-blue?style=flat-square&logo=telegram)](https://t.me/FlClash)

基于ClashMeta的多平台代理客户端，简单易用，开源无广告。

## 关于本项目

这是一个基于 [FlClash](https://github.com/chen08209/FlClash)
修改的个人精简版，主要面向希望获得更纯粹代理体验的用户。

本版本在保留 FlClash 核心代理能力和 Material You 界面的基础上，
移除了 Firebase 等非必要集成，并对部分功能与界面进行了精简，
使软件更加专注于日常代理使用。

### 主要特点

- **精简体验**：减少非必要功能与第三方服务，保留常用的代理管理能力。
- **自动选点**：可为指定策略组自动测试节点延迟，并优先从台湾、新加坡和日本节点中选择当前延迟较低的节点。
- **自动维护**：软件运行期间会定期重新检测节点；当前节点不可用或延迟过高时，也会自动尝试切换。
- **开源无广告**：继续遵循 GNU GPL v3 许可证开放源代码。

> [!NOTE]
> 本项目是 FlClash 的个人修改版，并非 FlClash 官方发行版本。
> 该版本于 2026 年 7 月开始修改，完整变更请查阅 Git 提交记录。

on Desktop:
<p style="text-align: center;">
    <img alt="desktop" src="snapshots/desktop.gif">
</p>

on Mobile:
<p style="text-align: center;">
    <img alt="mobile" src="snapshots/mobile.gif">
</p>

## Features

✈️ 多平台: Android, Windows, macOS and Linux

💻 自适应多个屏幕尺寸,多种颜色主题可供选择

💡 基本 Material You 设计, 类[Surfboard](https://github.com/getsurfboard/surfboard)用户界面

☁️ 支持通过WebDAV同步数据

✨ 支持一键导入订阅, 深色模式

## Use

### Linux

⚠️ 使用前请确保安装以下依赖

   ```bash
    sudo apt-get install libayatana-appindicator3-dev
    sudo apt-get install libkeybinder-3.0-dev
   ```

### Android

支持下列操作

   ```bash
    com.fyy6658.sipclash.action.START
    
    com.fyy6658.sipclash.action.STOP
    
    com.fyy6658.sipclash.action.TOGGLE
   ```

## Download

<a href="https://github.com/fyy6658-arch/FlClash/releases"><img alt="在 GitHub 下载 sipclash" src="snapshots/get-it-on-github.svg" width="200px"/></a>

### Homebrew

```bash
brew tap chen08209/tap
brew install --cask flclash
```

## Build

1. 更新 submodules
   ```bash
   git submodule update --init --recursive
   ```

2. 安装 `Flutter` 以及 `Golang` 环境

3. 构建应用

    - android

        1. 安装  `Android SDK` ,  `Android NDK`

        2. 设置 `ANDROID_NDK` 环境变量

        3. 运行构建脚本

           ```bash
           dart setup.dart android
           ```

    - windows

        1. 你需要一个windows客户端

        2. 安装 `GCC`，`Inno Setup`

        3. 运行构建脚本

           ```bash
           dart setup.dart windows
           ```

    - linux

        1. 你需要一个linux客户端

        2. 依赖会由 setup 脚本自动安装，也可以手动安装：
           ```bash
           sudo apt-get install -y libayatana-appindicator3-dev libkeybinder-3.0-dev
           ```

        3. 运行构建脚本

           ```bash
           dart setup.dart linux
           ```

    - macOS

        1. 你需要一个macOS客户端

        2. 运行构建脚本

           ```bash
           dart setup.dart macos
           ```

## Star

支持开发者的最简单方式是点击页面顶部的星标（⭐）。

<p style="text-align: center;">
    <a href="https://api.star-history.com/svg?repos=chen08209/FlClash&Date">
        <img alt="start" width=50% src="https://api.star-history.com/svg?repos=chen08209/FlClash&Date"/>
    </a>
</p>
