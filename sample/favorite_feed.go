package main

import (
	"fmt"

	"github.com/yukihir0/hbapi"
)

func main() {
	user := "yukihir0"
	params := hbapi.NewHBFavoriteFeedParams(user)
	//params.SetUser("yukihir0")
	//params.SetPage(1)
	//params.SetUntil(time.Now().Add(7 * 24 * -time.Hour))
	//params.EnableWithMe()
	//params.DisableWithMe()

	feed, err := hbapi.GetFavoriteFeed(params)
	if err != nil {
		panic(err)
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
