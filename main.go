//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	ap "github.com/moemoe89/simple-implementation-jwt-golang/api"
	"github.com/moemoe89/simple-implementation-jwt-golang/routers"

	"github.com/kataras/iris/v12"
)

// main will run the application
func main() {
	svc := ap.NewService()
	app := routers.GetRouter(svc)
	app.Run(iris.Addr(":8789"))
}
