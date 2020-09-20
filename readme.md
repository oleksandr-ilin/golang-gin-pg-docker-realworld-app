# ![RealWorld Example App](logo.png)

[![Build Status](https://travis-ci.org/wangzitian0/golang-gin-starter-kit.svg?branch=master)](https://travis-ci.org/wangzitian0/golang-gin-starter-kit)
[![codecov](https://codecov.io/gh/wangzitian0/golang-gin-starter-kit/branch/master/graph/badge.svg)](https://codecov.io/gh/wangzitian0/golang-gin-starter-kit)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/wangzitian0/golang-gin-starter-kit/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/wangzitian0/golang-gin-starter-kit?status.svg)](https://godoc.org/github.com/wangzitian0/golang-gin-starter-kit)

> ### Golang/Gin codebase containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.

This codebase was created to demonstrate a fully fledged fullstack application built with **Golang/Gin** including CRUD operations, authentication, routing, pagination, and more.
This is PostgreSQL version with Docker support.

# How it works

```
.
├── hello.go
├── common
│   ├── utils.go        //small tools function
│   └── database.go     //DB connect manager
├── users
│   ├── models.go       //data models define & DB operation
│   ├── serializers.go  //response computing & format
│   ├── routers.go      //business logic & router binding
│   ├── middlewares.go  //put the before & after logic of handle request
│   └── validators.go   //form/json checker
└── articles
    ├── models.go       //data models define & DB operation
    ├── serializers.go  //response computing & format
    ├── routers.go      //business logic & router binding
    └── validators.go   //form/json checker
```

# Getting started

## Install the Golang

https://golang.org/doc/install

## Environment Config

If you will clone this project into the `/Users/username/goprojects/src`
make sure your ~/.\*shrc have those varible:

```
➜  echo $GOPATH
/Users/username/goprojects/
➜  echo $GOROOT
/usr/local/go/
➜  echo $PATH
...:/usr/local/go/bin:/Users/username/goprojects/:/bin:/usr/local/go//bin
```

## Install Fresh

I used Fresh can help build without reload

https://github.com/pilu/fresh

```
cd
go get -u github.com/pilu/fresh
go get -u golang.org/x/crypto/...
```

## Start

```
➜  go run hello.go
➜  fresh
```

## Testing

```
➜  go test -v ./... -cover
```

## Docker

To build docker image:

```bash
docker build -f Dockerfile -t go/real-world-app .
```

To run container:
```bash
`docker run -p 8080:8080 --add-host pg-host:192.168.0.107 -it image_name`
```