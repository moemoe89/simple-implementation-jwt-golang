//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	"simple-implementation-jwt-golang/routers"

	"github.com/kataras/iris/v12"
)

func main() {
	app := routers.GetRouter()
	app.Run(iris.Addr(":8787"))
}
