PORT ?= 8080

.PHONY: build build-dev start start-dev stop

build:
	docker build -f Dockerfile -t fc-retrieval-provider .

build-dev:
	go build -v cmd/provider/main.go

start:
	docker-compose up

start-dev:
	go run cmd/provider/main.go --host 0.0.0.0 --port $(PORT)

stop:
	docker-compose stop
