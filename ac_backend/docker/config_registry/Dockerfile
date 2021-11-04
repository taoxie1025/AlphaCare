FROM golang:1.14 AS BUILDER

# Install protobuf compiler
RUN apt update && apt install -y protobuf-compiler && apt install -y golang-goprotobuf-dev

# copy proto definition (all protoc files are needed)
ADD /proto/. /go/src/alphacare/ac_backend/proto/
ADD /third_party/. /go/src/alphacare/ac_backend/third_party/

# generate go server code
WORKDIR /go/src/alphacare/ac_backend/
RUN for PROTO_FILE in proto/*.proto; do protoc --proto_path=proto --proto_path=./third_party --go_out=plugins=grpc:proto ${PROTO_FILE}; done

# copy files
ADD /logging/. /go/src/alphacare/ac_backend/logging/
ADD /services/config_registry/. /go/src/alphacare/ac_backend/services/config_registry/
ADD go.mod /go/src/alphacare/ac_backend/

# apply config
RUN mv /go/src/alphacare/ac_backend/services/config_registry/cmd/configs/dev.yaml /go/src/alphacare/ac_backend/services/config_registry/cmd/configs/config.yaml

# download dependencies
WORKDIR /go/src/alphacare/ac_backend/
RUN go mod tidy

# build executable
WORKDIR /go/src/alphacare/ac_backend/services/config_registry/cmd/
RUN go build -o main .

FROM golang:1.14
COPY --from=BUILDER /go/src/alphacare/ac_backend/services/config_registry/cmd/ .

EXPOSE 50052
CMD ["./main"]