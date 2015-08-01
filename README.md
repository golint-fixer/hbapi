# hbapi [![Build Status](https://travis-ci.org/yukihir0/hbapi.svg?branch=master)](https://travis-ci.org/yukihir0/hbapi) [![Coverage Status](https://coveralls.io/repos/yukihir0/hbapi/badge.svg?branch=master)](https://coveralls.io/r/yukihir0/hbapi?branch=master)

"hbapi" is library for hatena bookmark api.

## Install

```
go get github.com/yukihir0/hbapi
```

## How to use

### Use official API

#### Hatena Bookmark Entry Information API
```
url := "..."
entry, _ := hbapi.GetEntryInfo(url)
fmt.Println(entry.Count)
```

#### Hatena Bookmark Count API
```
// single url
url := "..."
count, _ := hbapi.GetBookmarkCount(url)
fmt.Printf("%s : %d\n", url, count)

// mutiple url
urls := []string{"...", "...", "..."}
counts, _ := hbapi.GetBookmarkCounts(urls)
for url, count := range counts {
  fmt.Printf("%s : %d\n", url, count)
}
```

#### Hatena Bookmark Feed
```
user := "..."
params := hbapi.NewFeedParams(user)
feed, _ := hbapi.GetFeed(params)
for _, item := range feed.Items {
	fmt.Println(item.Title)
}
```

#### Hatena Bookmark Favorite Feed
```
user := "..."
params := hbapi.NewFavoriteFeedParams(user)
feed, _ := hbapi.GetFavoriteFeed(params)
for _, item := range feed.Items {
	fmt.Println(item.Title)
}
```

### Use unofficial API

#### Hatena Bookmark Hot Entry Feed
```
feed, _ := hbapi.GetHotEntryFeed()
for _, item := range feed.Items {
  fmt.Println(item.Title)
}
```

## License

Copyright &copy; 2015 yukihir0
