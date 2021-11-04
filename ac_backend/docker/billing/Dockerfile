FROM golang:1.14 AS BUILDER

# Install protobuf compiler
RUN apt update && apt install -y protobuf-compiler && apt install -y golang-goprotobuf-dev

# copy proto definition
ADD /proto/billing.proto /go/src/alphacare/ac_backend/proto/
ADD /third_party/. /go/src/alphacare/ac_backend/third_party/

# generate go server code
WORKDIR /go/src/alphacare/ac_backend/
RUN for PROTO_FILE in proto/*.proto; do protoc --proto_path=proto --proto_path=./third_party --go_out=plugins=grpc:proto ${PROTO_FILE}; done

# copy files
ADD /logging/. /go/src/alphacare/ac_backend/logging/
ADD /services/billing/. /go/src/alphacare/ac_backend/services/billing/
ADD go.mod /go/src/alphacare/ac_backend/

# apply config
RUN mv /go/src/alphacare/ac_backend/services/billing/cmd/configs/dev.yaml /go/src/alphacare/ac_backend/services/billing/cmd/configs/config.yaml

# download dependencies
WORKDIR /go/src/alphacare/ac_backend/
RUN go mod tidy

# build executable
WORKDIR /go/src/alphacare/ac_backend/services/billing/cmd/
RUN go build -o main .

FROM golang:1.14
COPY --from=BUILDER /go/src/alphacare/ac_backend/services/billing/cmd/ .

EXPOSE 50057
CMD ["./main"]