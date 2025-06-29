.PHONY: help
.PHONY: build run clean test
.PHONY: swagger
.PHONY: podman-build podman-run

help:
	@echo "Usage: make [target]"
	@echo "  build        : ビルド"
	@echo "  run          : ローカル実行 (ビルド後)"
	@echo "                 例: PORT=8081 HELLO_MESSAGE=Hi make run"
	@echo "  clean        : バイナリ削除"
	@echo "  test         : Goテスト実行"
	@echo "  swagger      : Swaggerドキュメント生成"
	@echo "  podman-build : Podmanイメージビルド"
	@echo "  podman-run   : Podmanでコンテナ実行"
	@echo "                 例: make podman-run PORT=8081 HELLO_MESSAGE=Hi"

build:
	go mod tidy && go build -o ./build/http/main ./cmd/http/main.go

run: build
	PORT=$(PORT) HELLO_MESSAGE=$(HELLO_MESSAGE) ./build/http/main

clean:
	rm -f ./build/http/main

test:
	go test ./...

swagger:
	swag init -g cmd/http/main.go -o ./docs

podman-build:
	podman build -t hello-world-service .

podman-run: podman-build
	podman run --rm -e PORT=$(PORT) -e HELLO_MESSAGE="$(HELLO_MESSAGE)" -p $(or $(PORT),8081):$(or $(PORT),8081) hello-world-service

podman-ps:
	podman ps

podman-prune:
	podman system prune -a -f
