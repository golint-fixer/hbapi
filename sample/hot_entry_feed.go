package main

import (
	"fmt"

	"github.com/yukihir0/hbapi"
)

func main() {
	feed, err := hbapi.GetHotEntryFeed()
	if err != nil {
		panic(err)
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}
}
