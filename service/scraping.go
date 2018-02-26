package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/konojunya/goroutine-sample/model"
)

var (
	requestURL = "https://twitter.com/intent/user?screen_name="
)

// GetUserFromTwitter twitter.comからユーザーの情報を取得する
func GetUserFromTwitter(id string) (model.User, error) {
	fmt.Println("scraping start at:" + id)
	doc, err := goquery.NewDocument(requestURL + id)
	if err != nil {
		return model.User{}, err
	}

	el := doc.Find("h2 a")
	screenName, ok := el.Attr("href")
	if !ok {
		return model.User{}, errors.New("Can't read screen_name")
	}

	name := strings.TrimSpace(el.Text())
	icon, ok := el.Find("img").Attr("src")
	if !ok {
		return model.User{}, errors.New("Can't read profile image url")
	}

	description := strings.TrimSpace(doc.Find(".note").Text())

	fmt.Println("scraping end at:" + id)

	return model.User{
		ID:                   id,
		ScreeName:            screenName,
		Name:                 name,
		Description:          description,
		ProfileImageURLHttps: icon,
	}, nil

}
