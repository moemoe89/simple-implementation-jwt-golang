//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package routers

import (
	ap "simple-implementation-jwt-golang/api"

	"github.com/kataras/iris/v12"
)

func GetRouter() *iris.Application {
	router := iris.New()
	router.Get("/ping", ap.Ping)

	ctrl := ap.NewCtrl()

	router.Get("/generate", ctrl.Generate)
	router.Get("/parse", ctrl.Parse)

	return router
}