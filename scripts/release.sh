#!/bin/bash

set -e

# 检查是否提供了版本号
if [[ -z "$1" ]]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.0.0"
    exit 1
fi

VERSION=$1

echo "Creating release for version $VERSION"

# 更新版本号
sed -i "s/go 1.25/go 1.25/" go.mod

# 运行测试
echo "Running tests..."
go test -v ./...

# 构建所有平台
echo "Building binaries..."
./scripts/build.sh release

# 创建发布包
RELEASE_DIR="release"
mkdir -p $RELEASE_DIR

# 打包每个平台的二进制文件
for file in build/*; do
    if [[ -f "$file" ]]; then
        filename=$(basename "$file")
        if [[ "$filename" != "$PROJECT_NAME" ]]; then
            tar -czf "$RELEASE_DIR/${filename}.tar.gz" -C build "$filename"
        fi
    fi
done

echo "Release packages created in $RELEASE_DIR/"
echo "You can now create a GitHub release and upload these packages."