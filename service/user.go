package service

import (
	"time"

	"github.com/konojunya/go-routine-sample/model"
)

func GetUser(id int) model.User {
	time.Sleep(3 * time.Second)

	user := model.User{
		ID:   id,
		Name: "jun",
		Age:  20,
	}

	return user
}
