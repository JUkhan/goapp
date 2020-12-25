package service

import "log"

type (
	AuthService interface {
		Login(username, password string) bool
	}
	authService struct {
		authorizedUserName, authorizedPassword string
	}
)

func NewLoginService() AuthService {
	return &authService{
		authorizedUserName: "jukhan",
		authorizedPassword: "test",
	}
}
func (service *authService) Login(username, password string) bool {
	log.Println("DATA:;;", username, password)
	return service.authorizedUserName == username &&
		service.authorizedPassword == password
}
