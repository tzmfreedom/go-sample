.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up -d

.PHONY: console
console:
	docker-compose exec fuse /bin/bash

.PHONY: unmount
unmount:
	-@fusermount -u /tmp/x

.PHONY: run
run: unmount
	go run main.go
