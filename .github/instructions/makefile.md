# Makefile構成方針

## ルートMakefileの役割
- **オーケストレーション専用**
    - 複数マイクロサービスの一括ビルド・起動・停止などを `podman-compose` で管理
    - 個別サービスのビルドや実行は呼び出さない
- 例: `all-podman-compose-build`, `all-podman-compose-up`, `all-podman-compose-down` など

## 各マイクロサービスのMakefileの役割
- **各サービスの開発・ビルド・テスト・ローカル実行は各ディレクトリ内で閉じる**
    - 例: `build`, `run`, `test`, `clean` など
- サービスごとに独立したMakefileを配置
- サービス追加・削除時もルートMakefileの修正は不要

---

## サンプル構成

```
/Makefile                # ルート: orchestration (podman-compose) のみ
/microservices/hello-world/Makefile   # サービス個別の開発用
/microservices/hello2/Makefile        # サービス個別の開発用
```

### ルートMakefile（例）
```makefile
.PHONY: help all-podman-compose-build all-podman-compose-up all-podman-compose-down

help:
	@echo "Usage: make [target]"
	@echo "  all-podman-compose-build : 全サービスのイメージビルド"
	@echo "  all-podman-compose-up    : 全サービスをバックグラウンド起動"
	@echo "  all-podman-compose-down  : 全サービス停止・削除"

all-podman-compose-build:
	podman-compose -f podman-compose.yaml build

all-podman-compose-up:
	podman-compose -f podman-compose.yaml up -d

all-podman-compose-down:
	podman-compose -f podman-compose.yaml down
```

### 各マイクロサービスのMakefile（例: microservices/hello-world/Makefile）
```makefile
.PHONY: build run clean test

build:
	go build -o main main.go

run: build
	./main

clean:
	rm -f main

test:
	go test ./...
```

---

## メリット
- サービスごとの開発・CI/CDが独立しやすい
- ルートMakefileはシンプルに保てる
- サービス追加・削除時の影響範囲が小さい

---

## 補足
- サービスごとのMakefileは自由に拡張可能（lint, test, build, deploy など）
- ルートMakefileはpodman-composeやk8s用のオーケストレーション専用にするのが推奨
