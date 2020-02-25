//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	"simple-implementation-jwt-golang/routers"

	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestPingRoute(t *testing.T) {
	router := routers.GetRouter()

	e := httptest.New(t, router)
	e.GET("/ping").Expect().Status(httptest.StatusOK)
}
