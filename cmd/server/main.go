package main

import (
	controller "go-learn/controllers"
	middleware "go-learn/middlewares"
	service "go-learn/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {
	server := gin.Default()

	server.Use(
		gin.Recovery(),
		//* basic auth
		// middleware.BasicAuthorization(),
		//* default gin logger
		gin.Logger(),
		//* custom logger
		// middleware.Logger(),
		//* this used to get more information about the request
		//* like the headers, the body, the query params, etc
		// gindump.Dump(),
	)

	apiRoutes := server.Group("/api")
	{
		apiRoutes.POST("/login", func(ctx *gin.Context) {
			token, err := loginController.Login(ctx)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"token": token})
			}
		})

		videoRoutes := apiRoutes.Group("/video", middleware.AuthorizeJWT())
		{
			videoRoutes.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, videoController.FindAll())
			})

			videoRoutes.POST("/", func(ctx *gin.Context) {
				err := videoController.Save(ctx)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					ctx.JSON(http.StatusOK, gin.H{"message": "video created successfully"})
				}
			})
		}
	}

	post := os.Getenv("PORT")
	if post == "" {
		post = "localhost:8082"
	} else {
		post = ":8082"
	}
	server.Run(post)
}
