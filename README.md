# System Monitor

一个用于Deepin Linux系统的轻量级系统监控工具，支持温度和功耗监控，并提供可拖动的悬浮窗口界面。

## 功能特性

- 实时监控CPU和GPU温度
- 监控CPU和内存使用率
- 估算系统功耗
- 悬浮窗口界面
- 支持窗口拖动
- 低系统资源占用

## 系统要求

- Deepin 23.1 或其他基于Debian的Linux发行版
- Go 1.25 或更高版本

## 安装依赖

```bash
# 安装必要的系统工具
sudo apt update
sudo apt install lm-sensors

# 配置传感器（可选，用于获取真实硬件温度）
sudo sensors-detect
编译和运行
# 克隆项目
git clone <repository-url>
cd system-monitor

# 初始化Go模块
go mod tidy

# 编译
go build -o system-monitor

# 运行
./system-monitor