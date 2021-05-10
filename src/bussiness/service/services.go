package service

import (
	"github.com/yayanbachtiar/music-chart/src/bussiness/domain"
	"github.com/yayanbachtiar/music-chart/src/bussiness/service/users"
)

type Services struct {
	user users.UserItf
}

func Init(dom *domain.Domain) *Services {
	return &Services{
		user: users.InitUserServices(dom.User),
	}
}
