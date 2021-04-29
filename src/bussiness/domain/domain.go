package domain

import (
	"github.com/yayanbachtiar/music-chart/src/bussiness/domain/users"
)

type Domain struct {
	User  users.UserItf
}

func InitDomain() *Domain {
	return &Domain{
		User: users.InitUserDom(),
	}
}
