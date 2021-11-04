#!/bin/bash

cd ac_backend
export PATH="$PATH:$(go env GOPATH)/bin"
go install github.com/golang/protobuf/protoc-gen-go

for PROTO_FILE in proto/*.proto; do protoc --proto_path=proto --proto_path=./third_party --go_out=plugins=grpc:proto ${PROTO_FILE}; done
