//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"time"

	"github.com/kataras/iris/v12"
)

var starTime = time.Now()

// Ping will handle the ping endpoint
func Ping(c iris.Context) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := starTime.In(loc)

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(res)
}
