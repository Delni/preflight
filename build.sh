#!/bin/bash

OUT="bin/"
TARGETS_OS="darwin freebsd linux openbsd windows"
TARGETS_ARCH="amd64 arm64"
EXT=""

for os in $TARGETS_OS; do
    for arch in $TARGETS_ARCH; do
        echo "Building $arch on $os"
        ext=$([[ "$os" == "windows" ]] && echo ".exe" || echo "")
        GOOS=$os GOARCH=$arch go build -o "${OUT}preflight${ext}"
        zip -r "${OUT}preflight_${os}_${arch}.zip" "${OUT}preflight${ext}"
        rm "${OUT}preflight${ext}"
    done
done