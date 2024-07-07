package service

import entity "go-learn/entities"

// interface to define the methods of the service
type VideoService interface {
	FindAll() []entity.Video
	Save(entity.Video)
}

// struct to implement the interface
type videoService struct {
	videos []entity.Video // slice of videos
}

// constructor function to create a new instance of the interface
// and return the implementation of the interface
func NewVideoService() VideoService {
	// return the implementation of the interface as a pointer of the struct
	return &videoService{
		videos: make([]entity.Video, 0), // Initialize videos slice
	}
}

// implement the methods of the interface
// should specify which struct you use that implements the interface
func (service *videoService) Save(video entity.Video) {
	service.videos = append(service.videos, video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
