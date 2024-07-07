package repository

import (
	entity "go-learn/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video) error
	Delete(video entity.Video) error
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, errr := gorm.Open("sqlite3", "test.db")
	if errr != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("failed to close database")
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) error {
	// update the video
	err := db.connection.Save(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *database) Delete(video entity.Video) error {
	err := db.connection.Delete(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	// Set("gorm:auto_preload", true) is used to automatically load the associated data (Autor) when querying the Video table
	return videos
}
