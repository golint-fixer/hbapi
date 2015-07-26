package main

import (
	"fmt"

	"github.com/yukihir0/hbapi"
)

func main() {
	user := "yukihir0"
	params := hbapi.NewHBFeedParams(user)
	params.Page = 0
	//params.Tag = "golang"
	//params.Date = "20150725"
	//params.URL = "http://d.hatena.ne.jp"

	feed, err := hbapi.GetFeed(params)
	if err != nil {
		panic(err)
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
