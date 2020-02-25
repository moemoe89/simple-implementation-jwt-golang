[![Build Status](https://travis-ci.org/moemoe89/simple-implementation-jwt-golang.svg?branch=master)](https://travis-ci.org/moemoe89/simple-implementation-jwt-golang)
[![Coverage Status](https://coveralls.io/repos/github/moemoe89/simple-implementation-jwt-golang/badge.svg?branch=master)](https://coveralls.io/github/moemoe89/simple-implementation-jwt-golang?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/simple-implementation-jwt-golang)](https://goreportcard.com/report/github.com/moemoe89/simple-implementation-jwt-golang)

# Simple-Implementation-JWT-Golang

![demo](simple-implementation-jwt-golang-demo.gif)

Simple Implementation JWT Using Golang (Iris Framework) as Programming Language

This repository created to show the example for using [JWT](https://jwt.io/) with this library [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ simple-implementation-jwt-golang/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> simple-implementation-jwt-golang
```

# Demo
Use this host [https://simple-jwt-golang.herokuapp.com/](https://simple-jwt-golang.herokuapp.com/)

# Running & Using Application
Make sure you already have [POSTMAN](https://www.getpostman.com/)

	    > Running the application using `go run main.go`
	    > Open your POSTMAN.
		> Fill the url with localhost:8787/generate and choose GET method
		> You will get the response below :
		  {
            "message": "Successfully generated token.",
            "success": true,
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkX3RpbWUiOjE0Nzk0NDExNTEsInRoaXJkX3BhcnR5IjoibW9tbyJ9.P5sZwzCJD1DYptpuCp4hIyY5pGGnOB7m6ZHHi_mBEi4"
          }
        > For checking the token, fill the url with localhost:3000/parse and choose GET method. Fill the Header variable with Authorization and the value with token that you've get. You will get the response below :
          {
            "data": {
              "third_party": "momo",
              "expired_time": 1479441229
            },
            "message": "Successfully parsed token.",
            "success": true
          }

# Running with Docker
```
$ docker build -t simple-implementation-jwt-golang .
$ docker run -d -p 8787:8787 --name simple-jwt-go simple-implementation-jwt-golang
```
How to stop the docker
```
$ docker ps
```
Copy the CONTAINER ID
```$xslt
$ docker stop {$CONTAINER_ID}
```