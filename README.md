# Simple-Implementation-JWT-Golang

![demo](simple-implementation-jwt-golang-demo.gif)

This repository created to show the example for using [JWT](https://jwt.io/) with this library [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
		
# Running & Using Application
Make sure you already have [POSTMAN](https://www.getpostman.com/)

	    > Running the application using `go run main.go`
	    > Open your POSTMAN.
		> Fill the url with localhost:3000/generate and choose GET method
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
