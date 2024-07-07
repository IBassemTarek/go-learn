package api

import (
	controller "go-learn/controllers"
	"go-learn/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoAPI interface {
	GetVideos(ctx *gin.Context)
	CreateVideos(ctx *gin.Context)
	UpdateVideos(ctx *gin.Context)
	DeleteVideos(ctx *gin.Context)
}

type videoAPI struct {
	videoController controller.VideoController
}

func NewVideoAPI(videoController controller.VideoController) VideoAPI {
	return &videoAPI{
		videoController: videoController,
	}
}

// paths information

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
func (api *videoAPI) GetVideos(ctx *gin.Context) {
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
func (api *videoAPI) CreateVideos(ctx *gin.Context) {
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
func (api *videoAPI) UpdateVideos(ctx *gin.Context) {
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
func (api *videoAPI) DeleteVideos(ctx *gin.Context) {
	err := api.videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{Message: err.Error()})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{Message: "video deleted successfully"})
	}
}
