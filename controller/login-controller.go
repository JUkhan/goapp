package controller

import (
	"github.com/JUkhan/goapp/entity"
	"github.com/JUkhan/goapp/service"
	"github.com/gin-gonic/gin"
)

type (
	LoginController interface {
		Login(*gin.Context) string
	}
	loginController struct {
		loginService service.LoginService
		jwtService   service.JWTService
	}
)

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}
func (con *loginController) Login(c *gin.Context) string {
	var credentials entity.Credential

	err := c.ShouldBindJSON(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := con.loginService.Login(credentials.UserName, credentials.Password)
	if isAuthenticated {
		return con.jwtService.GenerateToken(credentials.UserName, true)
	}
	return ""
}
