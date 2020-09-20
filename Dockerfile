# build stage
FROM golang:alpine AS build-env
RUN set -eux; \
    apk add --no-cache --virtual .build-deps \
    git gcc libc-dev; \
    go get -u github.com/pilu/fresh; \
    go get -u golang.org/x/crypto/... ;
# apk del .build-deps;
ENV GOPATH /go
ENV GOROOT /usr/local/go
WORKDIR /go/src/golang-gin-pg-docker-realworld-app
COPY . /go/src/golang-gin-pg-docker-realworld-app
RUN go build hello.go

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/golang-gin-pg-docker-realworld-app/hello /app/
EXPOSE 8080
ENTRYPOINT ./hello