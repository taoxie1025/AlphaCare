FROM golang:1.14 AS BUILDER

# Install protobuf compiler
RUN apt update && apt install -y protobuf-compiler && apt install -y golang-goprotobuf-dev

# copy proto definition
ADD /proto/*.proto /go/src/alphacare/ac_backend/proto/
ADD /third_party/. /go/src/alphacare/ac_backend/third_party/

# generate go server code
WORKDIR /go/src/alphacare/ac_backend/
RUN for PROTO_FILE in proto/*.proto; do protoc --proto_path=proto --proto_path=./third_party --go_out=plugins=grpc:proto ${PROTO_FILE}; done

# copy files
ADD /logging/. /go/src/alphacare/ac_backend/logging/
ADD /services/user_registry/. /go/src/alphacare/ac_backend/services/user_registry/
ADD go.mod /go/src/alphacare/ac_backend/

# apply config
RUN mv /go/src/alphacare/ac_backend/services/user_registry/cmd/configs/dev.yaml /go/src/alphacare/ac_backend/services/user_registry/cmd/configs/config.yaml

# download dependencies
WORKDIR /go/src/alphacare/ac_backend/
RUN go mod tidy

# build executable
WORKDIR /go/src/alphacare/ac_backend/services/user_registry/cmd/
RUN go build -o main .

FROM golang:1.14
COPY --from=BUILDER /go/src/alphacare/ac_backend/services/user_registry/cmd/ .

EXPOSE 50051
CMD ["./main"]