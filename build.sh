#!/bin/bash

OUT="bin/"
TARGETS_OS="darwin freebsd linux openbsd windows"
TARGETS_ARCH="amd64 arm64"

for os in $TARGETS_OS; do
    for arch in $TARGETS_ARCH; do
        echo "Building $arch on $os"
        GOOS=$os GOARCH=$arch go build -o "${OUT}preflight_${os}_${arch}"
    done
done