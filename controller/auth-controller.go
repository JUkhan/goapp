package controller

import (
	"github.com/JUkhan/goapp/dto"
	"github.com/JUkhan/goapp/service"
	"github.com/gin-gonic/gin"
)

type (
	AuthController interface {
		Login(*gin.Context) string
	}
	authController struct {
		loginService service.AuthService
		jwtService   service.JWTService
	}
)

func NewAuthController(loginService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}
func (con *authController) Login(c *gin.Context) string {
	var credentials dto.Credentials

	err := c.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := con.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return con.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
