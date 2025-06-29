.PHONY: help
.PHONY: podman-compose-build podman-compose-up podman-compose-down podman-compose-ps

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "  podman-compose-build : 全サービスのイメージビルド"
	@echo "  podman-compose-up    : 全サービスをバックグラウンド起動"
	@echo "  podman-compose-down  : 全サービス停止・削除"

podman-compose-build:
	podman-compose -f podman-compose.yaml build

podman-compose-up:
	podman-compose -f podman-compose.yaml up -d

podman-compose-down:
	podman-compose -f podman-compose.yaml down

podman-compose-ps:
	podman-compose ps
