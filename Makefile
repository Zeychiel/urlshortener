.PHONY: build run test docker-build docker-run

build:
	go build -o app cmd/api/main.go

run: build
	./app

test:
	go test -v ./...

docker-build:
	docker build -t shortener .

docker-run: docker-build
	docker run shortener

up:
	docker-compose up --build

lint:
	golangci-lint run