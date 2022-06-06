FROM golang:1.18-alpine AS builder

WORKDIR /
ADD . .

WORKDIR /
RUN  CGO_ENABLED=0 go build -mod vendor -o /nedap-connector

FROM alpine:3.7

RUN apk -U add ca-certificates

WORKDIR /app

COPY --from=builder / /app/

EXPOSE 3000

ENTRYPOINT ./bookmark-man