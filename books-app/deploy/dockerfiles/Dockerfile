FROM golang:1.23.0 as builder

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make compile-grpc-server

FROM alpine:3.14.2
COPY --from=builder /service/build/out/grpc-books-server /grpc-books-server
COPY --from=builder /service/configs/grpc-books-server.yaml /grpc-books-server.yaml
ENTRYPOINT ["/grpc-books-server", "-configFile", "/grpc-books-server.yaml"]