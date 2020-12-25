package api

import (
	"net/http"

	"github.com/JUkhan/goapp/controller"
	"github.com/JUkhan/goapp/dto"
	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	loginController controller.AuthController
}

func NewAuthAPI(loginController controller.AuthController) *AuthApi {
	return &AuthApi{
		loginController: loginController,
	}
}

// Paths Information

// Authenticate godoc
// @Summary Provides a JSON Web Token
// @Description Authenticates a user and provides a JWT to Authorize API calls
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} dto.JWT
// @Failure 401 {object} dto.Response
// @Router /auth/token [post]
func (api *AuthApi) Authenticate(c *gin.Context) {
	token := api.loginController.Login(c)
	if token != "" {
		c.JSON(http.StatusOK, &dto.JWT{Token: token})
	} else {
		c.JSON(http.StatusUnauthorized, &dto.Response{Message: "Not Authorized"})
	}
}
