package api

import (
	controller "go-learn/controllers"
	"go-learn/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppAPI interface {
	Authenticate(ctx *gin.Context)
	GetVideos(ctx *gin.Context)
	CreateVideos(ctx *gin.Context)
	UpdateVideos(ctx *gin.Context)
	DeleteVideos(ctx *gin.Context)
}

type appAPI struct {
	loginController controller.LoginController
	videoController controller.VideoController
}

func NewAppAPI(loginController controller.LoginController, videoController controller.VideoController) AppAPI {
	return &appAPI{
		loginController: loginController,
		videoController: videoController,
	}
}

// paths information

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
func (api *appAPI) Authenticate(ctx *gin.Context) {
	token, err := api.loginController.Login(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{Message: err.Error()})
	} else if token == "" {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{Message: "Invalid credentials"})
	} else {
		ctx.JSON(http.StatusOK, &dto.Token{Token: token})
	}
}

// GetVideos godoc
// @Security Bearer
// @Summary List existing videos
// @Description get all videos
// @ID list-videos
// @Tags Videos
// @Accept json
// @Produce json
// @Success 200 {array} entity.Video
// @Failure 401 {object} dto.Response
// @Router /video [get]
func (api *appAPI) GetVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.videoController.FindAll())
}

// CreateVideos godoc
// @Security Bearer
// @Summary create a new video
// @Description create a new video
// @ID create-video
// @Tags Videos
// @Accept json
// @Produce json
// @Param video body entity.Video true "Video"
// @Success 200 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /video [post]
func (api *appAPI) CreateVideos(ctx *gin.Context) {
	err := api.videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{Message: "video created successfully"})
	}
}

// UpdateVideos godoc
// @Security Bearer
// @Summary update a video
// @Description update a video by ID
// @ID update-video
// @Tags Videos
// @Accept json
// @Produce json
// @Param video body entity.Video true "Video"
// @Param id path int true "Video ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /video/{id} [put]
func (api *appAPI) UpdateVideos(ctx *gin.Context) {
	err := api.videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{Message: "video updated successfully"})
	}
}

// DeleteVideos godoc
// @Security Bearer
// @Summary Delete a video
// @Description Delete a video by ID
// @ID delete-video
// @Tags Videos
// @Accept json
// @Produce json
// @Param id path int true "Video ID"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /video/{id} [delete]
func (api *appAPI) DeleteVideos(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{Message: "video deleted successfully"})
	}
}
