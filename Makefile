.PHONY: build build-dev start stop

build:
  docker build -f Dockerfile -t fc-retrieval-provider .

build-dev:
  go build -v cmd/provider/main.go

start:
  docker-compose up

start-dev:
  go run cmd/provider/main.go

stop:
  docker-compose stop
