package main

import (
	"go-learn/api"
	controller "go-learn/controllers"
	docs "go-learn/docs"
	middleware "go-learn/middlewares"
	"go-learn/repository"
	service "go-learn/services"
	"os"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.NewVideoService(videoRepository)
	videoController controller.VideoController = controller.NewVideoController(videoService)

	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJWTService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @SecurityRequirement Bearer
func main() {

	host := os.Getenv("PORT")
	if host == "" {
		host = "localhost:8082"
	}

	// swagger 2.0 meta information
	docs.SwaggerInfo.Title = "Demo Video API"
	docs.SwaggerInfo.Description = "This is a simple video API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// defer the close of the db connection
	defer videoRepository.CloseDB()

	server := gin.New()
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

	videoAPI := api.NewVideoAPI(videoController)
	authAPI := api.NewAuthAPI(loginController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/login", authAPI.Authenticate)
		}
		videoRoutes := apiRoutes.Group("/video", middleware.AuthorizeJWT())
		{
			videoRoutes.GET("/", videoAPI.GetVideos)
			videoRoutes.POST("/", videoAPI.CreateVideos)
			videoRoutes.PUT("/:id", videoAPI.UpdateVideos)
			videoRoutes.DELETE("/:id", videoAPI.DeleteVideos)
		}
	}
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.Run(host)
}
