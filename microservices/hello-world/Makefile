.PHONY: help
.PHONY: build run clean test
.PHONY: podman-build podman-run

help:
	@echo "Usage: make [target]"
	@echo "  build        : ビルド"
	@echo "  run          : ローカル実行 (ビルド後)"
	@echo "  clean        : バイナリ削除"
	@echo "  test         : Goテスト実行"
	@echo "  podman-build : Podmanイメージビルド"
	@echo "  podman-run   : Podmanでコンテナ実行"

build:
	go mod tidy && go build -o main main.go

run: build
	./main

clean:
	rm -f main

test:
	go test ./...

podman-build:
	podman build -t hello-world-service .

podman-run: podman-build
	podman run --rm -p 8080:8080 hello-world-service
