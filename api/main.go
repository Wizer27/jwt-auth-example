package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserData struct {
	UserId int `json:"user_id"`
}

func auth_middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-API-KEY")

		if token != "secret" {
			c.AbortWithStatusJSON(401, gin.H{
				"error": "unathorized",
			})
			return
		}

		c.Next()

	}
}

func get_current_user_from_jwt(token_str string) float64 {
	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte("secret"), nil
	})

	if err != nil {
		fmt.Println("Ошибка:", err)
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id, ok2 := claims["user_id"].(float64)
		if ok2 {
			return user_id
		} else {
			return 0
		}
	}
	return 0
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
			return
		}

		if user_data.UserId == 0 {
			c.JSON(400, gin.H{
				"error": "invalid user_id",
			})
			return
		}
		token := generate_jwt_token(user_data.UserId)

		c.JSON(200, gin.H{
			"token": token,
		})

	})

	r.GET("/protected", auth_middleware(), func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(401, gin.H{
				"error": "invalid token",
			})
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		user_id := get_current_user_from_jwt(token)
		if user_id == 0 {
			c.JSON(401, gin.H{
				"error": "invalid jwt token",
			})
			return
		}

		user_id_int := int(user_id)
		fmt.Println("User ", user_id_int, "logged in")

		c.JSON(200, gin.H{
			"message": "welcome",
		})

	})
	r.Run(":8080")
}
