
# Deepin Thermometer

[![Release](https://github.com/yourusername/system-monitor/actions/workflows/release.yml/badge.svg)](https://github.com/yourusername/system-monitor/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/system-monitor)](https://goreportcard.com/report/github.com/yourusername/system-monitor)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

一个轻量级的系统监控工具，支持Linux、macOS和Windows平台，提供实时的CPU温度、GPU温度、使用率和功耗监控。

## 🌟 功能特性

- 🌡️ **实时温度监控** - 监控CPU和GPU温度
- 📊 **系统使用率** - 实时显示CPU和内存使用率  
- ⚡ **功耗估算** - 估算系统当前功耗
- 🖱️ **悬浮窗口** - 无边框可拖动的悬浮监控面板
- 🚀 **低资源占用** - 轻量级设计，不影响系统性能
- 🖥️ **跨平台支持** - 支持Linux、macOS和Windows
- 🎯 **简单操作** - 左键拖动，右键关闭，双击退出

## 📋 系统要求

### Linux系统
- **支持的发行版**: Ubuntu 18.04+, Debian 10+, CentOS 7+, Fedora 30+, Deepin 20+
- **必要依赖**: 
  - `libgl1-mesa-dev`
  - `xorg-dev` 
  - `libgtk-3-dev`
  - `build-essential`

### macOS系统
- **最低版本**: macOS 10.12+
- **推荐**: 最新版本以获得最佳性能

### Windows系统
- **最低版本**: Windows 7+
- **推荐**: Windows 10/11

## 🚀 安装方式

### 📦 二进制安装（推荐）

从[Releases](https://github.com/yourusername/system-monitor/releases)页面下载适合您系统的预编译二进制文件：

1. 访问 [GitHub Releases](https://github.com/yourusername/system-monitor/releases)
2. 下载对应平台的压缩包
3. 解压文件
4. 直接运行可执行文件

### 🔧 源码编译安装

#### 克隆仓库
```bash
git clone https://github.com/yourusername/system-monitor.git
cd deepin-thermometer
```

#### 安装系统依赖
```bash
# Linux (Ubuntu/Debian/Deepin)
sudo apt update
sudo apt install libgl1-mesa-dev xorg-dev libgtk-3-dev build-essential

# Linux (CentOS/RHEL/Fedora)  
sudo yum install mesa-libGL-devel libX11-devel gtk3-devel gcc

# Linux (Arch Linux)
sudo pacman -S mesa libx11 gtk3 gcc
```

#### 构建项目
```bash
# 安装Go依赖
go mod tidy

# 构建
go build -o deepin-thermometer ./cmd/deepin-thermometer/

# 或使用Makefile
make build
```

## ▶️ 使用方法

### 基本运行
```bash
# 直接运行
./deepin-thermometer

# 或者使用Makefile
make run
```

### 操作指南
- **启动程序**: 双击可执行文件或在终端运行
- **移动窗口**: 鼠标左键按住窗口任意位置拖动
- **关闭程序**:
    - 鼠标右键点击窗口
    - 双击窗口
    - 按 `Ctrl+C` (终端运行时)
- **查看信息**: 窗口实时显示系统监控数据

### 显示信息说明
```
CPU: 45.2°C (25.6%)     # CPU温度和使用率
GPU: 42.1°C             # GPU温度
MEM: 48.3%              # 内存使用率
POWER: 22.8W            # 估算功耗
```

## 🛠️ 开发指南

### 项目结构
```
system-monitor/
├── cmd/                    # 主程序入口
│   └── system-monitor/     # 主程序文件
├── internal/               # 内部包
│   ├── monitor/            # 系统监控核心逻辑
│   │   ├── monitor.go      # 监控器主逻辑
│   │   └── sensors.go      # 传感器读取
│   └── ui/                 # 用户界面
│       └── window.go       # Ebiten窗口实现
├── scripts/                # 自动化脚本
│   ├── build.sh            # 构建脚本
│   ├── install-deps.sh     # 依赖安装脚本
│   └── release.sh          # 发布脚本
├── .github/workflows/      # GitHub Actions配置
│   └── release.yml         # 自动发布工作流
├── Makefile                # Make构建文件
├── go.mod                  # Go模块文件
└── go.sum                  # Go依赖校验和
```

### 构建脚本

#### 使用Makefile（推荐）
```bash
# 构建当前平台版本
make build

# 运行程序
make run

# 安装系统依赖
make install-deps

# 清理构建文件
make clean

# 构建发布版本
make release VERSION=v1.0.0
```

#### 手动构建
```bash
# 构建当前平台
go build -o deepin-thermometer ./cmd/deepin-thermometer/

# 构建Linux版本
GOOS=linux GOARCH=amd64 go build -o deepin-thermometer-linux ./cmd/deepin-thermometer/

# 构建Windows版本
GOOS=windows GOARCH=amd64 go build -o deepin-thermometer.exe ./cmd/deepin-thermometer/

# 构建macOS版本
GOOS=darwin GOARCH=amd64 go build -o deepin-thermometer-macos ./cmd/deepin-thermometer/
```

### 脚本说明

#### build.sh - 构建脚本
```bash
# 构建当前平台
./scripts/build.sh

# 构建发布版本
./scripts/build.sh release
```

#### install-deps.sh - 依赖安装脚本
```bash
# 自动检测系统并安装依赖
./scripts/install-deps.sh
```

#### release.sh - 发布脚本
```bash
# 创建指定版本的发布包
./scripts/release.sh v1.0.0
```

## 🔧 配置说明

### 传感器配置
程序会自动检测系统硬件传感器。如果没有找到真实传感器，将使用模拟数据。

#### Linux传感器配置
```bash
# 安装lm-sensors
sudo apt install lm-sensors

# 配置传感器
sudo sensors-detect

# 测试传感器
sensors
```

### 环境变量
```bash
# 设置窗口初始位置
export WINDOW_X=1000
export WINDOW_Y=50

# 设置更新频率（秒）
export UPDATE_INTERVAL=1
```

## 🐛 故障排除

### 常见问题

#### Linux图形错误
```bash
# 错误信息: vulkan/vulkan.h: No such file or directory
sudo apt install libvulkan-dev

# 错误信息: libGL相关错误
sudo apt install libgl1-mesa-dev xorg-dev
```

#### 传感器无法读取
```bash
# 检查是否有传感器
ls /sys/class/hwmon/

# 手动安装传感器工具
sudo apt install lm-sensors
sudo sensors-detect
```

#### 窗口显示异常
```bash
# 检查桌面环境支持
echo $XDG_CURRENT_DESKTOP

# 尝试不同的窗口管理器设置
export GDK_BACKEND=x11
./deepin-thermometer
```

### 日志调试
```bash
# 启用详细日志
export DEBUG=1
./deepin-thermometer

# 查看系统日志
journalctl -f
```

## 🤝 贡献指南

欢迎提交Issue和Pull Request！

### 开发流程
1. Fork项目到您的GitHub账户
2. 创建功能分支：`git checkout -b feature/new-feature`
3. 提交更改：`git commit -am 'Add new feature'`
4. 推送分支：`git push origin feature/new-feature`
5. 创建Pull Request

### 代码规范
- 遵循Go语言编码规范
- 添加必要的注释
- 编写单元测试
- 保持代码简洁易读

### 测试
```bash
# 运行所有测试
go test -v ./...

# 运行覆盖率测试
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📦 发布流程

### GitHub Actions自动发布
1. 在GitHub上创建Release并打tag
2. GitHub Actions会自动构建所有平台版本
3. 自动生成Release并上传二进制文件

### 手动发布
```bash
# 创建版本tag
git tag v1.0.0
git push origin v1.0.0

# 构建发布包
./scripts/release.sh v1.0.0

# 上传到GitHub Release
```

## 📄 许可证

MIT License

Copyright (c) 2024 Your Name

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

## 👤 作者

**Your Name**
- GitHub: [@yourusername](https://github.com/yourusername)
- Email: your.email@example.com

## 🙏 致谢

- [Ebiten](https://ebiten.org/) - 优秀的2D游戏引擎
- [Go语言](https://golang.org/) - 强大的编程语言
- 所有贡献者和支持者

---
*Made with ❤️ using Go and Ebiten*
```