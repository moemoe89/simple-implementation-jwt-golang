//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	"github.com/moemoe89/simple-implementation-jwt-golang/api/api_struct/model"
	"github.com/moemoe89/simple-implementation-jwt-golang/routers"
	"github.com/moemoe89/simple-implementation-jwt-golang/api/mocks"

	"errors"
	"testing"

	"github.com/kataras/iris/v12/httptest"
	"net/http"
)

func TestSuccessGenerate(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Generate").Return("foo", 0, nil)

	router := routers.GetRouter(mockService)

	e := httptest.New(t, router)
	e.GET("/generate").Expect().Status(httptest.StatusOK)
}

func TestFailGenerate(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Generate").Return("", http.StatusInternalServerError, errors.New("Error"))

	router := routers.GetRouter(mockService)

	e := httptest.New(t, router)
	e.GET("/generate").Expect().Status(httptest.StatusInternalServerError)
}

func TestSuccessParse(t *testing.T) {
	var jwtModel *model.JWTModel
	mockService := new(mocks.Service)
	mockService.On("Parse", "foo").Return(jwtModel, 0, nil)

	router := routers.GetRouter(mockService)

	e := httptest.New(t, router)
	e.GET("/parse").WithHeader("Authorization", "foo").Expect().Status(httptest.StatusOK)
}

func TestFailParse(t *testing.T) {
	mockService := new(mocks.Service)
	mockService.On("Parse", "foo").Return(nil, http.StatusInternalServerError, errors.New("Error"))

	router := routers.GetRouter(mockService)

	e := httptest.New(t, router)
	e.GET("/parse").WithHeader("Authorization", "foo").Expect().Status(httptest.StatusInternalServerError)
}
