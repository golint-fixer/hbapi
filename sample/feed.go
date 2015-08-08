package main

import (
	"fmt"
	"time"

	"github.com/yukihir0/hbapi"
)

func main() {
	user := "yukihir0"
	params := hbapi.NewFeedParams(user)
	//params.SetUser("yukihir0")
	//params.SetPage(1)
	//params.SetTag("golang")
	params.SetDate(time.Date(2015, 7, 29, 0, 0, 0, 0, time.UTC))
	//params.SetURL("http://d.hatena.ne.jp")

	feed, err := hbapi.GetFeed(params)
	if err != nil {
		panic(err)
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
