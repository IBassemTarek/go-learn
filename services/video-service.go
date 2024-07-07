package service

import (
	entity "go-learn/entities"
	"go-learn/repository"
)

// interface to define the methods of the service
type VideoService interface {
	FindAll() []entity.Video
	Update(video entity.Video) error
	Delete(video entity.Video) error
	Save(entity.Video)
}

// struct to implement the interface
type videoService struct {
	videoRepository repository.VideoRepository
}

// constructor function to create a new instance of the interface
// and return the implementation of the interface
func NewVideoService(repo repository.VideoRepository) VideoService {
	// return the implementation of the interface as a pointer of the struct
	return &videoService{
		videoRepository: repo,
	}
}

// implement the methods of the interface
// should specify which struct you use that implements the interface
func (service *videoService) Save(video entity.Video) {
	service.videoRepository.Save(video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video) error {
	return service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) error {
	return service.videoRepository.Delete(video)
}
