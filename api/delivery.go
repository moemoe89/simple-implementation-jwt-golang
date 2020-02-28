//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"github.com/moemoe89/simple-implementation-jwt-golang/api/api_struct/model"
	cons "github.com/moemoe89/simple-implementation-jwt-golang/constant"

	"github.com/kataras/iris/v12"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
	svc Service
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(svc Service) *ctrl {
	return &ctrl{svc: svc}
}

func (ct *ctrl) Generate(c iris.Context) {
	tokenString, status, err := ct.svc.Generate()
	if err != nil {
		c.StatusCode(status)
		c.JSON(model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	resp := map[string]interface{}{
		"token": tokenString,
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(model.NewGenericResponse(iris.StatusOK, cons.ERR, []string{"Successfully generated token."}, resp))
}

func (ct *ctrl) Parse(c iris.Context) {
	unparsedToken := c.GetHeader("Authorization")

	tokenParsing, status, err := ct.svc.Parse(unparsedToken)
	if err != nil {
		c.StatusCode(status)
		c.JSON(model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}

	resp := map[string]interface{}{
		"data": tokenParsing,
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(model.NewGenericResponse(iris.StatusOK, cons.ERR, []string{"Successfully parsed token."}, resp))
}