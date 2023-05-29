# syntax = docker/dockerfile:1.3
FROM --platform=$BUILDPLATFORM golang:1.20.3-buster as builder

WORKDIR /app
COPY ./go.mod ./go.sum ./

COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH task build:core

FROM scratch
WORKDIR /app

COPY --from=builder /app/core /app

ENTRYPOINT ["./core"]
