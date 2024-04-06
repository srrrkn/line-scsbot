## ビルド

コンテナビルド時のコマンド

```bash
DOCKER_BUILDKIT=1 docker build . -f build/Dockerfile -t line-scs-bot:latest
```

## 起動
コンテナ起動時のコマンド

```bash
docker compose -f deployments/docker-compose.yml up -d
```

## Goのinit
```
go mod init main.go
go mod tidy
go fmt
```