package main

import (
	"fmt"
	"github.com/yukihir0/hbapi"
)

func main() {
	// single url
	url := "http://developer.hatena.ne.jp"
	count, err := hbapi.GetBookmarkCount(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s : %d\n", url, count)

	// mutliple
	urls := []string{"https://github.com", "https://bitbucket.org", "http://stackoverflow.com"}
	counts, err := hbapi.GetBookmarkCounts(urls)
	if err != nil {
		panic(err)
	}
	for url, count := range counts {
		fmt.Printf("%s : %d\n", url, count)
	}
}
