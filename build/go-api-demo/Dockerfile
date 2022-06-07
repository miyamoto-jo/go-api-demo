FROM golang:1.16.15-alpine as builder

WORKDIR /build
COPY ../../ ./
RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./

FROM alpine
COPY --from=builder /build/go-api-demo /opt/app/
ENTRYPOINT ["/opt/app/go-api-demo"]