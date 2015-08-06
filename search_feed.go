package hbapi

import (
	"time"
)

// SearchFeed represents hatena bookmark search feed.
type SearchFeed struct {
	Title       string
	Link        string
	Description string
	Items       []SearchFeedItem
}

// SearchFeedItem represents hatena bookmark search feed item.
type SearchFeedItem struct {
	Title         string
	Link          string
	Description   string
	Content       string
	Creator       string
	Date          time.Time
	BookmarkCount int
	Subject       []string
}

func newSearchFeed(f UnifiedFeed) SearchFeed {
	ret := SearchFeed{
		Title:       f.Title,
		Link:        f.Link,
		Description: f.Description,
		Items:       newSearchFeedItems(f.Items),
	}
	return ret
}

func newSearchFeedItems(items []UnifiedFeedItem) []SearchFeedItem {
	ret := []SearchFeedItem{}
	for _, i := range items {
		ret = append(ret, SearchFeedItem{
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
