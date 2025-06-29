# microservices/hello-world サービス要件

- HTTP GET / で "Hello, world!" を返すGo製マイクロサービスであること
- シングルバイナリで動作し、8080番ポートでリッスンすること
- コンテナ化されていること（Dockerfileあり）
- サービス個別のMakefileでビルド・実行・テスト・Podman操作ができること
