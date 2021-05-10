package users

import (
	"github.com/yayanbachtiar/music-chart/src/bussiness/domain/users"
	"github.com/yayanbachtiar/music-chart/src/bussiness/model"
)

type UserService struct {
	userDom users.UserItf
}

func (u *UserService) GetUsers() []model.User {
	panic("implement me")
}

func (u *UserService) SaveUser() []model.User {
	panic("implement me")
}

type UserItf interface {
	GetUsers()[]model.User
	SaveUser()[]model.User
}

func InitUserServices(user users.UserItf) UserItf {
	return &UserService{
		userDom: user,
	}
}
