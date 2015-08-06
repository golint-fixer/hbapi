package hbapi

import (
	"time"
)

// HotEntryFeed represents hatena bookmark hot entry feed.
type HotEntryFeed struct {
	Title       string
	Link        string
	Description string
	Items       []HotEntryFeedItem
}

// HotEntryFeedItem represents hatena bookmark hot entry feed item.
type HotEntryFeedItem struct {
	Title         string
	Link          string
	Description   string
	Content       string
	Creator       string
	Date          time.Time
	BookmarkCount int
	Subject       []string
}

func newHotEntryFeed(f UnifiedFeed) HotEntryFeed {
	ret := HotEntryFeed{
		Title:       f.Title,
		Link:        f.Link,
		Description: f.Description,
		Items:       newHotEntryFeedItems(f.Items),
	}
	return ret
}

func newHotEntryFeedItems(items []UnifiedFeedItem) []HotEntryFeedItem {
	ret := []HotEntryFeedItem{}
	for _, i := range items {
		ret = append(ret, HotEntryFeedItem{
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
