## builder
FROM golang:1.17.5-alpine3.15 as migration-builder

WORKDIR /build
COPY . .
RUN cd go_task && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /build/main ./cmd/migration

## process
FROM alpine:3.15

COPY docker/migration /
COPY --from=migration-builder /build/main /main
RUN adduser -D task && chown task:task /main && chmod 755 /main

RUN apk update && \
    apk add mysql-client

USER task

CMD ["/start.sh"]
