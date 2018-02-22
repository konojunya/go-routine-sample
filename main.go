package main

import (
	"fmt"

	"github.com/konojunya/go-routine-sample/model"
	"github.com/konojunya/go-routine-sample/service"
)

var (
	ids   []int
	users []model.User
)

func init() {
	for i := 0; i < 1000; i++ {
		ids = append(ids, i)
	}
}

func main() {
	// アンチパターン
	// for _, id := range ids {
	// 	go func(id int) {
	// 		user := service.GetUser(id)
	// 		users = append(users, user)
	// 	}(id)
	// }

	// まっきsample
	userCh := make(chan model.User)

	for _, id := range ids {
		go func(id int, userCh chan model.User) {
			userCh <- service.GetUser(id)
		}(id, userCh)
	}

	for _ = range ids {
		user := <-userCh
		users = append(users, user)
	}
	fmt.Println(users)
}
