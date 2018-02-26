package service

import (
	"fmt"
)

var count = 0

// PostTwitterAPI twitter apiっぽい疑似アクセス
func PostTwitterAPI() ([]string, bool) {
	fmt.Println("twitter access!")
	count++
	if count > 10 {
		return nil, true
	}
	return []string{"konojunya", "makki0205", "kinokoruumu416"}, false
}
