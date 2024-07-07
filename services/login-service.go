package service

type LoginService interface {
	Login(userName, password string) bool
}

type loginService struct {
	// this as a demo we will use a hardcoded username and password
	adminUsername string
	adminPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		adminUsername: "admin",
		adminPassword: "admin",
	}
}

func (service *loginService) Login(userName, password string) bool {
	// this is a demo, in a real application you would check the username and password against a database
	return service.adminUsername == userName && service.adminPassword == password
}
