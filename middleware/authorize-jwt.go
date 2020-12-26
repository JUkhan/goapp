package middleware

import (
	"log"
	"net/http"

	"github.com/JUkhan/goapp/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(c *gin.Context) {
	const BearerSchema = "Bearer "
	authorizeHead := c.GetHeader("Authorization")
	if authorizeHead == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	tokenString := authorizeHead[len(BearerSchema):]
	token, err := service.NewJWTService().ValidateToken(tokenString)

	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		log.Println(claims)
	} else {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// func sss() {
// 	authHeader := c.GetHeader("Authorization")
// 	tokenString := authHeader[len(BEARER_SCHEMA):]

// 	token, err := service.NewJWTService().ValidateToken(tokenString)

// 	if token.Valid {
// 		claims := token.Claims.(jwt.MapClaims)
// 		log.Println("Claims[Name]: ", claims["name"])
// 		log.Println("Claims[Admin]: ", claims["admin"])
// 		log.Println("Claims[Issuer]: ", claims["iss"])
// 		log.Println("Claims[IssuedAt]: ", claims["iat"])
// 		log.Println("Claims[ExpiresAt]: ", claims["exp"])
// 	} else {
// 		log.Println(err)
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 	}
// }
