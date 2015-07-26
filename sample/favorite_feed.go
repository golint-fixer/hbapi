package main

import (
	"fmt"

	"github.com/yukihir0/hbapi"
)

func main() {
	user := "yukihir0"
	params := hbapi.NewHBFavoriteFeedParams(user)
	params.Page = 0
	//params.Until = time.Now().Add(7 * 24 * -time.Hour)
	//params.WithMe = true

	feed, err := hbapi.GetFavoriteFeed(params)
	if err != nil {
		panic(err)
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
