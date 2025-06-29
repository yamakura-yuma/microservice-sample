#!/bin/bash

echo "[INFO] Listing shared libraries in the container..."

echo "[INFO] go"
echo "go version: $(go version)"

echo "[INFO] node"
echo "node version: $(node --version)"
echo "npm version: $(npm --version)"

echo "[INFO] podman"
echo "$(podman --version)"