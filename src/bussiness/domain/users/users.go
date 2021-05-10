package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/yayanbachtiar/music-chart/src/bussiness/model"
	"io/ioutil"
	"os"
)

type user struct {
	sql *sql.DB
}

func (u *user) SaveUser() []model.User {
	panic("implement me")
}

type UserItf interface {
	GetUsers()[]model.User
	SaveUser()[]model.User
}

func InitUserDom(sql *sql.DB) *user {
	return &user{sql: sql}
}

func (u *user) GetUsers() []model.User {
	// Open our jsonFile
	jsonFile, err := os.Open("db/user.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	var users []model.User
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil
	}
	return users
}

