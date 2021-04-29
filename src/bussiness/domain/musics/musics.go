package users

import (
	"encoding/json"
	"fmt"
	"github.com/yayanbachtiar/music-chart/src/bussiness/model"
	"io/ioutil"
	"os"
)

type music struct {

}

func (u *music) SaveMusic() []model.User {
	panic("implement me")
}

type MusicsItf interface {
	GetMusic()[]model.User
	SaveMusic()[]model.User
}

func InitUserDom() *music {
	return &music{}
}

func (u *music) GetMusic() []model.User {
	// Open our jsonFile
	jsonFile, err := os.Open("db/music.json")
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

