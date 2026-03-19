package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserData struct {
	UserId int `json:"user_id"`
}

func auth_middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token != "secret" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unathorized",
			})
			return
		}

		c.Next()

	}
}

func generate_jwt_token(user_id int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
	})
	token_string, _ := token.SignedString([]byte("secret"))
	return token_string
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.POST("/get_token", auth_middleware(), func(c *gin.Context) {

		var user_data UserData

		err := c.BindJSON(&user_data)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid json",
			})
		}

		token := generate_jwt_token(user_data.UserId)

		c.JSON(200, gin.H{
			"token": token,
		})

	})

	r.Run(":8080")
}
