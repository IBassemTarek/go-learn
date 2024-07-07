package api

import (
	controller "go-learn/controllers"
	"go-learn/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthAPI interface {
	Authenticate(ctx *gin.Context)
}

type authAPI struct {
	loginController controller.LoginController
}

func NewAuthAPI(loginController controller.LoginController) AuthAPI {
	return &authAPI{
		loginController: loginController,
	}
}

// Authenticate godoc
// @Summary Provides a JWT token
// @Description authenticate with username and password to get the JWT token
// @ID authentication
// @Consume application/json
// @Param credentials body dto.Credentials true "Credentials"
// @Produce json
// @Tags Auth
// @Success 200 {object} dto.Token
// @Failure 401 {object} dto.Response
// @Router /auth/login [post]
func (api *authAPI) Authenticate(ctx *gin.Context) {
	token, err := api.loginController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{Message: err.Error()})
	} else if token == "" {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{Message: "Invalid credentials"})
	} else {
		ctx.JSON(http.StatusOK, &dto.Token{Token: token})
	}
}
