//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	"github.com/moemoe89/simple-implementation-jwt-golang/routers"

	"github.com/kataras/iris/v12"
)

// main will run the application
func main() {
	app := routers.GetRouter()
	app.Run(iris.Addr(":8787"))
}
