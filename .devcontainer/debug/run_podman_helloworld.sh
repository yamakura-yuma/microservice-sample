
#!/bin/bash
# podmanでhello-worldイメージを実行するスクリプト

set -e

echo "[INFO] podman version:"
podman --version

echo "[INFO] Running hello-world container with podman..."
podman run --rm hello-world

if [ $? -eq 0 ]; then
  echo "[SUCCESS] hello-worldコンテナの実行に成功しました。"
else
  echo "[ERROR] hello-worldコンテナの実行に失敗しました。"
  exit 1
fi
