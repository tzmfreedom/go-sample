.PHONY: proto/build
proto/build:
	protoc -I helloworld helloworld/hello.proto --go-grpc_out=helloworld

.PHONY: client/run
client/run:
	go run client/main.go

.PHONY: server/run
server/run:
	go run server/main.go
