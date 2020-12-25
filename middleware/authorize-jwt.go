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
	tokenString := authorizeHead[len(BearerSchema):]
	token, err := service.NewJWTService().ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		log.Println(claims)
	} else {
		log.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
