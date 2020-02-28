//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package api_test

import (
	ap "github.com/moemoe89/simple-implementation-jwt-golang/api"
	"github.com/moemoe89/simple-implementation-jwt-golang/routers"

	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestPingRoute(t *testing.T) {
	svc := ap.NewService()
	router := routers.GetRouter(svc)

	e := httptest.New(t, router)
	e.GET("/ping").Expect().Status(httptest.StatusOK)
}
