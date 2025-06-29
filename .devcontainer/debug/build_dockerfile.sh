#!/bin/bash
# podmanでdevcontainer用Dockerfileがビルドできるかテストするスクリプト

set -e

# ビルドコンテキストとDockerfileのパス
DOCKERFILE="${CODESPACE_VSCODE_FOLDER}/.devcontainer/Dockerfile"
IMAGE_NAME="devcontainer-podman-test:latest"

echo "[INFO] podman version:"
podman --version

echo "[INFO] Building devcontainer image with podman..."
podman build \
    -t "$IMAGE_NAME" \
    -f "$DOCKERFILE" \
    "$CONTEXT_DIR"

if [ $? -eq 0 ]; then
  echo "[SUCCESS] devcontainerイメージのビルドに成功しました。"
else
  echo "[ERROR] devcontainerイメージのビルドに失敗しました。"
  exit 1
fi
