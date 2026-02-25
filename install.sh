#!/bin/bash

OS=$(uname -s)
if [ "$OS" != "Linux" ]; then
    echo "This script only supports Linux/Android platforms."
    exit 1
fi

ARCH=$(uname -m)
case "$ARCH" in
    aarch64|arm64) ARCH="arm64" ;;
    armv7*|armv8*) ARCH="arm" ;;
    x86_64)        ARCH="amd64" ;;
    i386|i686)     ARCH="386" ;;
    *)             echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac

BINARY="ClashFa-Wizard"
ARCHIVE="${BINARY}-${OS}-${ARCH}.tar.gz"
LATEST_VERSION=$(curl -fsSL https://raw.githubusercontent.com/10ium/worker--Wizard/main/VERSION)

if [ -x "./${BINARY}" ]; then
    INSTALLED_VERSION=$("./${BINARY}" --version)
    echo "Installed version: $INSTALLED_VERSION"
    echo "Latest version: ${LATEST_VERSION}"

    if [ "${INSTALLED_VERSION}" = "${LATEST_VERSION}" ]; then
        echo "Wizard is up to date. Running..."
        exec ./"${BINARY}"
    else
        echo "Updating to version ${LATEST_VERSION}..."
    fi
else
    echo "Wizard not found on device. Installing version ${LATEST_VERSION}..."
fi

echo "Downloading ${ARCHIVE}..."
curl -L -# -o "${ARCHIVE}" "https://github.com/10ium/worker--Wizard/releases/latest/download/${ARCHIVE}" && \
tar xzf "${ARCHIVE}" && \
chmod +x "./${BINARY}" && \
exec ./"${BINARY}"