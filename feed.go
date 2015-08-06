package hbapi

import (
	"time"
)

// Feed represents hatena bookmark feed.
type Feed struct {
	Title        string
	Link         string
	Description  string
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Items        []FeedItem
}

// FeedItem represents hatena bookmark feed item.
type FeedItem struct {
	Title         string
	Link          string
	Description   string
	Content       string
	Creator       string
	Date          time.Time
	BookmarkCount int
	Subject       []string
}

func newFeed(f UnifiedFeed) Feed {
	ret := Feed{
		Title:        f.Title,
		Link:         f.Link,
		Description:  f.Description,
		StartIndex:   f.StartIndex,
		ItemsPerPage: f.ItemsPerPage,
		TotalResults: f.TotalResults,
		Items:        newFeedItems(f.Items),
	}
	return ret
}

func newFeedItems(items []UnifiedFeedItem) []FeedItem {
	ret := []FeedItem{}
	for _, i := range items {
		ret = append(ret, FeedItem{
			Title:         i.Title,
			Link:          i.Link,
			Description:   i.Description,
			Content:       i.Content,
			Creator:       i.Creator,
			Date:          i.Date,
			BookmarkCount: i.BookmarkCount,
			Subject:       i.Subject,
		})
	}
	return ret
}
