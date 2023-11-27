# Golang Gin REST API Sample

* based on the official sample
  - https://go.dev/doc/tutorial/web-service-gin

## Quick Start

```bash
go version
# go version go1.21.3 darwin/arm64

make test
python -m http.server --directory ./test_results/ 3001
# http://localhost:3001/cover.html

go run ./main.go
# or
./run_server.sh

curl http://localhost:8080/
# {"msg":"hello world"}
curl http://localhost:8080/albums
curl http://localhost:8080/albums/2

# or build and run
go build ./main.go && ./main
```

## Initial Project Setup

```bash
go mod init example/web-service-gin
```

## Setup goimports staticcheck

* https://pkg.go.dev/golang.org/x/tools/cmd/goimports
* https://staticcheck.dev/docs/getting-started/

```bash
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest

# in .zshrc
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

## Docker

* Dockerfile official sample
  - https://docs.docker.com/language/golang/build-images/

```bash
docker --version
# Docker version 24.0.6, build ed223bc

docker-compose --version
# Docker Compose version v2.23.0-desktop.1

docker-compose build
docker-compose up -d
docker-compose ps

curl http://localhost:8080/
# {"msg":"hello world"}

docker-compose exec app sh

# clean up
docker-compose down
```

## Reference

* Golang Web API開発入門ガイド
  - https://qiita.com/MYu/items/1c1a93fba22035b45d6a
* Go言語で基本的なCRUD操作を行うREST APIを作成
  - https://dev.classmethod.jp/articles/go-sample-rest-api/
* 【golang】DockerでGoの開発環境を構築する
  - https://qiita.com/___yusuke49/items/0f6577b0b6af5f63671b
