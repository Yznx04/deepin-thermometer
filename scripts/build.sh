#!/bin/bash

set -e

# 设置变量
PROJECT_NAME="system-monitor"
VERSION=${1:-"dev"}
BUILD_DIR="build"
PLATFORMS=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

echo "Building $PROJECT_NAME version $VERSION"

# 清理之前的构建
rm -rf $BUILD_DIR
mkdir -p $BUILD_DIR

# 构建当前平台版本
echo "Building for current platform..."
go build -o $BUILD_DIR/$PROJECT_NAME -ldflags "-X main.version=$VERSION" ./cmd/system-monitor/

# 构建多平台版本
if [[ "$1" == "release" ]]; then
    echo "Building for multiple platforms..."
    for PLATFORM in "${PLATFORMS[@]}"; do
        GOOS=${PLATFORM%/*}
        GOARCH=${PLATFORM#*/}

        OUTPUT_NAME="$PROJECT_NAME-$GOOS-$GOARCH"
        if [[ $GOOS == "windows" ]]; then
            OUTPUT_NAME+=".exe"
        fi

        echo "Building $OUTPUT_NAME..."
        env GOOS=$GOOS GOARCH=$GOARCH go build -o $BUILD_DIR/$OUTPUT_NAME -ldflags "-X main.version=$VERSION" ./cmd/system-monitor/
    done
fi

echo "Build completed! Output files are in $BUILD_DIR/"