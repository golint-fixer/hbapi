package main

import (
	"fmt"
	"github.com/yukihir0/hbapi"
)

func main() {
	url := "http://developer.hatena.ne.jp"
	entry, err := hbapi.GetEntryInfo(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", entry)
}
