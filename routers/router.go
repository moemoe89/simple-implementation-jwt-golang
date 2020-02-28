//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/simple-implementation-jwt-golang/api"

	"github.com/kataras/iris/v12"
)

// GetRouter will create a variable that represent the iris.Application
func GetRouter(svc ap.Service) *iris.Application {
	router := iris.New()
	router.Get("/ping", ap.Ping)

	ctrl := ap.NewCtrl(svc)

	router.Get("/generate", ctrl.Generate)
	router.Get("/parse", ctrl.Parse)

	return router
}