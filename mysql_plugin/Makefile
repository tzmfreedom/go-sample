.PHONY: build
build: format
	docker run --rm -v $(PWD):/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 docker_go go build -buildmode=c-shared -o main.so main.go

.PHONY: format
format:
	gofmt -w .

.PHONY: docker/build
docker/build:
	docker build -f Dockerfile.build . -t docker_go

.PHONY: docker/mysql
docker/mysql: build
	docker build . -t my_mysql

.PHONY: docker/mysql/run
docker/mysql/run: docker/mysql
	docker run -e MYSQL_ROOT_PASSWORD=mysql -d -p 3306:3306 my_mysql
