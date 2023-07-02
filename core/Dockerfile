# syntax = docker/dockerfile:1.3
FROM --platform=$BUILDPLATFORM golang:1.20.5-buster as builder

RUN go install github.com/go-task/task/v3/cmd/task@v3.27.1

WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod download

COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH task core:build

FROM alpine:3.18.2
WORKDIR /app

COPY --from=builder /app/core/bin/vizimind /app

ENTRYPOINT ["./vizimind"]
