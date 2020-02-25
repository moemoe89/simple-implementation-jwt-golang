//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"simple-implementation-jwt-golang/api/api_struct/model"
	cons "simple-implementation-jwt-golang/constant"

	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

const(
	hmacSecret = "momo12345"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl() *ctrl {
	return &ctrl{}
}

func (ct *ctrl) Generate(c iris.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"third_party": "momo",
		"expired_time": time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSecret))

	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(model.NewGenericResponse(iris.StatusInternalServerError, cons.ERR, []string{err.Error()}, nil))
		return
	}

	resp := map[string]interface{}{
		"token":   tokenString,
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(model.NewGenericResponse(iris.StatusInternalServerError, cons.ERR, []string{"Successfully generated token."}, resp))
	return
}

func (ct *ctrl) Parse(c iris.Context) {
	unparsedToken := c.GetHeader("Authorization")

	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})

	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(model.NewGenericResponse(iris.StatusInternalServerError, cons.ERR, []string{err.Error()}, nil))
		return
	}

	raw, err := json.Marshal(token.Claims)

	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(model.NewGenericResponse(iris.StatusInternalServerError, cons.ERR, []string{err.Error()}, nil))
		return
	}

	var tokenParsing model.JWTModel
	err = json.Unmarshal(raw, &tokenParsing)

	if err != nil {
		c.StatusCode(iris.StatusInternalServerError)
		c.JSON(model.NewGenericResponse(iris.StatusInternalServerError, cons.ERR, []string{err.Error()}, nil))
		return

	} else {

		if time.Now().Unix() > int64(tokenParsing.ExpiredTime) {
			c.StatusCode(iris.StatusUnauthorized)
			c.JSON(model.NewGenericResponse(iris.StatusUnauthorized, cons.ERR, []string{"Token has been expired."}, nil))
			return
		}

	}

	resp := map[string]interface{}{
		"data": tokenParsing,
	}
	c.StatusCode(iris.StatusOK)
	c.JSON(model.NewGenericResponse(iris.StatusInternalServerError, cons.ERR, []string{"Successfully parsed token."}, resp))
	return
}