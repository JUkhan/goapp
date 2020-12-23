package service

type (
	LoginService interface {
		Login(username, password string) bool
	}
	loginService struct {
		authorizedUserName, authorizedPassword string
	}
)

func NewLoginService() LoginService {
	return &loginService{
		authorizedUserName: "jukhan",
		authorizedPassword: "test",
	}
}
func (service *loginService) Login(username, password string) bool {
	return service.authorizedUserName == username &&
		service.authorizedPassword == password
}
