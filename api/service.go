//
//  Simple JWT
//
//  Copyright © 2016. All rights reserved.
//

package api

import (
	"github.com/moemoe89/simple-implementation-jwt-golang/api/api_struct/model"

	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	hmacSecret = "momo12345"
)

// Service represent the services
type Service interface {
	Generate() (string, int, error)
	Parse(unparsedToken string) (*model.JWTModel, int, error)
}

type implService struct {
}

// NewService will create an object that represent the Service interface
func NewService() Service {
	return &implService{}
}

func (s *implService) Generate() (string, int, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"third_party":  "momo",
		"expired_time": time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return tokenString, 0, nil
}

func (s *implService) Parse(unparsedToken string) (*model.JWTModel, int, error) {
	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	raw, err := json.Marshal(token.Claims)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var tokenParsing model.JWTModel
	err = json.Unmarshal(raw, &tokenParsing)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if time.Now().Unix() > int64(tokenParsing.ExpiredTime) {
		return nil, http.StatusUnauthorized, errors.New("Token has been expired.")
	}

	return &tokenParsing, 0, nil
}