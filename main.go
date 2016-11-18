package main

import (
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"net/http"
	"time"
)

const(
	hmacSecret = "momo12345"
)

type JWTModel struct {
	ThirdParty string `json:"third_party"`
	ExpiredTime float64 `json:"expired_time"`
}

func GenerateToken(c *gin.Context) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"third_party": "momo",
		"expired_time": time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(hmacSecret))

	if err != nil {

		resp := gin.H{
			"success": false,
			"message": err.Error(),
		}

		c.IndentedJSON(http.StatusInternalServerError, resp)
		return

	} else {

		resp := gin.H{
			"success": true,
			"message": "Successfully generated token.",
			"token": tokenString,
		}

		c.IndentedJSON(http.StatusOK, resp)
		return
	}

}

func CheckToken(c *gin.Context) {

	unparsedToken := c.Request.Header.Get("Authorization")

	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})

	if err != nil {
		resp := gin.H{
			"success": false,
			"message": err.Error(),
		}

		c.IndentedJSON(http.StatusInternalServerError, resp)
		return

	}

	raw, err := json.Marshal(token.Claims)

	if err != nil {

		resp := gin.H{
			"success": false,
			"message": err.Error(),
		}

		c.IndentedJSON(http.StatusInternalServerError, resp)
		return

	}

	var tokenParsing *JWTModel
	err = json.Unmarshal(raw, &tokenParsing)

	if err != nil {

		resp := gin.H{
			"success": false,
			"message": err.Error(),
		}

		c.IndentedJSON(http.StatusInternalServerError, resp)
		return

	} else {

		if time.Now().Unix() > int64(tokenParsing.ExpiredTime) {

			resp := gin.H{
				"success": false,
				"message": "Token has been expired.",
			}

			c.IndentedJSON(http.StatusUnauthorized, resp)
			return

		} else {

			resp := gin.H{
				"success": true,
				"message": "Successfully parsed token.",
				"data": tokenParsing,
			}

			c.IndentedJSON(http.StatusOK, resp)
			return

		}

	}

}

func main() {

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))

	router.GET("/generate",GenerateToken)
	router.GET("/parse",CheckToken)
	router.Run(":3000")

}