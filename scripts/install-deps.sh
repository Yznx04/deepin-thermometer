#!/bin/bash

set -e

echo "Installing system dependencies..."

# 检测操作系统
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Ubuntu/Debian/Deepin
    if command -v apt &> /dev/null; then
        echo "Installing dependencies for Ubuntu/Debian/Deepin..."
        sudo apt update
        sudo apt install -y \
            libgl1-mesa-dev \
            xorg-dev \
            libgtk-3-dev \
            build-essential \
            git
    # CentOS/RHEL/Fedora
    elif command -v yum &> /dev/null; then
        echo "Installing dependencies for CentOS/RHEL/Fedora..."
        sudo yum install -y \
            mesa-libGL-devel \
            libX11-devel \
            gtk3-devel \
            gcc \
            git
    # Arch Linux
    elif command -v pacman &> /dev/null; then
        echo "Installing dependencies for Arch Linux..."
        sudo pacman -S --noconfirm \
            mesa \
            libx11 \
            gtk3 \
            gcc \
            git
    else
        echo "Unsupported package manager. Please install dependencies manually."
        exit 1
    fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Installing dependencies for macOS..."
    # macOS通常不需要额外的依赖
    if ! command -v brew &> /dev/null; then
        echo "Homebrew not found. Installing..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
    fi
else
    echo "Unsupported operating system. Please install dependencies manually."
    exit 1
fi

echo "System dependencies installed successfully!"