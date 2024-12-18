package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("Authorization")
		fmt.Println("TOKEN", tokenString)
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't get token"})
			return

		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
				return
			}
		}
		userId := claims["sub"]

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("id", userId)
		c.Next()

	}
}

/*func GetUserIdFromContext(ctx *gin.Context) float64 {
	tokenUserId, _ := ctx.Get("id")
	userId, _ := tokenUserId.(float64)
	return userId
}*/

func GetUserIdFromContext(ctx *gin.Context) int {
	tokenUserId, _ := ctx.Get("id")
	userId, _ := tokenUserId.(float64)
	return int(userId)
}
