.PHONY: proto/build
proto/build:
	protoc -I helloworld helloworld/hello.proto --go_out=plugins=grpc:helloworld

.PHONY: client/run
client/run:
	go run client/

.PHONY: server/run
server/run:
	go run server
