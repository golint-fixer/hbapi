package main

import (
	"fmt"

	"github.com/yukihir0/hbapi"
)

func main() {
	query := "golang"
	params := hbapi.NewSearchFeedParams(query)
	//params.SetQuery("api")
	//params.SetTarget2Text()
	//params.SetTarget2Tag()
	//params.SetTarget2Title()
	//params.SetSort2Recent()
	//params.SetSort2Popular()
	//params.SetUsers(100)
	//params.EnableSafe()
	//params.DisableSafe()
	//params.SetDateBegin(time.Date(2015, 7, 1, 0, 0, 0, 0, time.UTC))
	//params.SetDateEnd(time.Date(2015, 7, 15, 0, 0, 0, 0, time.UTC))

	feed, err := hbapi.GetSearchFeed(params)
	if err != nil {
		panic(err)
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
