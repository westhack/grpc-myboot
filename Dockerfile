FROM golang:alpine

WORKDIR /go/src/grpc-myboot
COPY . .

RUN go generate && go env && go build -o grpc-myboot .

FROM alpine:latest
LABEL MAINTAINER="452855489@qq.com"

WORKDIR /go/src/grpc-myboot

COPY --from=0 /go/src/grpc-myboot ./

EXPOSE 50051

ENTRYPOINT ./grpc-myboot -c config.docker.yaml
