.PHONY: protos

protos:
	@echo generating protos...
	@protoc -I=./:${GOPATH}/src/github.com/google/protobuf/src --go_out=plugins=grpc:./ ./messages.proto	
