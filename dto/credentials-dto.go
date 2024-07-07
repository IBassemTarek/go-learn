package dto

type Credentails struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
