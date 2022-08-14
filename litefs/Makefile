CONFIG=/app/litefs.yml

.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up -d

.PHONY: console
console:
	docker-compose exec litefs /bin/bash

consul:
	consul agent -dev

run:
	litefs -config $(CONFIG)

run/sqlite3:
	docker-compose exec litefs /bin/bash -c "sqlite3 $(FILE)"
