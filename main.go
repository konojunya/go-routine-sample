package main

import (
	"fmt"
	"sync"

	"github.com/konojunya/goroutine-sample/model"
	"github.com/konojunya/goroutine-sample/service"
)

func main() {
	// アンチパターン
	// for _, id := range ids {
	// 	go func(id int) {
	// 		user := service.GetUser(id)
	// 		users = append(users, user)
	// 	}(id)
	// }

	// まっきsample
	/*
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
	*/
	userCh := make(chan model.UserWithError)
	users := make([]model.User, 0)
	var wg sync.WaitGroup

	for {
		ids, next := service.PostTwitterAPI()
		for _, id := range ids {
			wg.Add(1)
			go func(id string) {
				user, err := service.GetUserFromTwitter(id)
				wg.Done()
				userCh <- model.UserWithError{
					User:  user,
					Error: err,
				}
			}(id)
		}

		for _ = range ids {
			res := <-userCh
			users = append(users, res.User)
		}

		if next {
			break
		}
	}
	wg.Wait()

	fmt.Println(users)

}
