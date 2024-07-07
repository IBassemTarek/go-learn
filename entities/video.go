package entity

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`    // btw 1 and 130
	Email     string `json:"email" binding:"required,email"` // required and email validation
}

type Video struct {
	// field and json serialization ( struct tags )
	Title       string `json:"title" binding:"min=2,max=100"` // min 2 and max 100
	Description string `json:"description" binding:"max=200"`
	URL         string `json:"url" binding:"required,url" validate:"is-not-google"` // required and url validation and custom validation
	Autor       Person `json:"autor" binding:"required"`
}
