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
func GetRouter() *iris.Application {
	router := iris.New()
	router.Get("/ping", ap.Ping)

	ctrl := ap.NewCtrl()

	router.Get("/generate", ctrl.Generate)
	router.Get("/parse", ctrl.Parse)

	return router
}