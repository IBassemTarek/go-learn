package entity

import "time"

type Person struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name" binding:"required" gorm:"type:varchar(32);not null"`
	LastName  string `json:"last_name" binding:"required" gorm:"type:varchar(32);not null"`
	Age       int16  `json:"age" binding:"gte=1,lte=130"`                                             // btw 1 and 130
	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256);UNIQUE;not null"` // required and email validation
}

type Video struct {
	// field and json serialization ( struct tags )
	ID          uint      `gorm:"primary_key;auto_increment" json:"id" swaggerignore:"true"`
	Title       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100);not null"`                       // min 2 and max 100
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(100);not null"`                       // max 200
	URL         string    `json:"url" binding:"required,url" validate:"is-not-google" gorm:"type:varchar(256);not null"` // required and url validation and custom validation
	Autor       Person    `json:"autor" binding:"required" gorm:"foreignkey:AutorID"`
	AutorID     uint64    `json:"-"` // non render this field
	CreatedAt   time.Time `json:"-" gorm:"CURRENT_TIMESTAMP;autoCreateTime" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"-" gorm:"CURRENT_TIMESTAMP;autoUpdateTime" swaggerignore:"true"`
}
