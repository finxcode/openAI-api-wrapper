# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE=on

WORKDIR /build

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../. ./
RUN go build -o chatgpt-wrapper main.go

FROM alpine

ENV OPENAI_API_KEY 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

WORKDIR /app

COPY --from=builder /build /app

CMD["./chatgpt-wrapper"]




