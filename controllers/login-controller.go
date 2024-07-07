package controller

import (
	"errors"
	dto "go-learn/dto"
	service "go-learn/services"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) (string, error)
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) (string, error) {
	var credentials dto.Credentails
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		return "", err
	}

	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jwtService.GenerateToken(credentials.Username, true), nil
	}
	return "", errors.New("authentication failed")
}
