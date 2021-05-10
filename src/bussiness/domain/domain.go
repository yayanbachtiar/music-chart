package domain

import (
	"database/sql"
	"github.com/yayanbachtiar/music-chart/src/bussiness/domain/users"
)

type Domain struct {
	User  users.UserItf
}

func InitDomain(sql *sql.DB) *Domain {
	return &Domain{
		User: users.InitUserDom(sql),
	}
}
