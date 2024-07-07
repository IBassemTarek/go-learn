package controller

import (
	entity "go-learn/entities"
	service "go-learn/services"
	"go-learn/validators"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

// take the service as a dependency
// to be able to call the service methods
type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func NewVideoController(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-not-google", validators.ValidateURL)
	return &videoController{
		service: service,
	}
}

func (controller *videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	// extract the json from the request conetxt (payload)
	// bind it to the video struct
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	// validate the video struct using the validate struct ( custom validation )
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}

func (controller *videoController) Update(ctx *gin.Context) error {
	var video entity.Video
	errBind := ctx.ShouldBindJSON(&video)
	if errBind != nil {
		return errBind
	}
	id, errParseId := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if errParseId != nil {
		return errParseId
	}
	video.ID = uint(id)
	err := validate.Struct(video)
	if err != nil {
		return err
	}
	updateErr := controller.service.Update(video)
	if updateErr != nil {
		return updateErr
	}
	return nil
}

func (controller *videoController) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, errParseId := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if errParseId != nil {
		return errParseId
	}
	video.ID = uint(id)
	deleteErr := controller.service.Delete(video)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}
