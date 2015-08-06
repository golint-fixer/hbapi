package hbapi

import (
	"time"
)

// FavoriteFeed represents hatena bookmark favorite feed.
type FavoriteFeed struct {
	Title       string
	Link        string
	Description string
	Items       []FavoriteFeedItem
}

// FavoriteFeedItem represents hatena bookmark favorite feed items.
type FavoriteFeedItem struct {
	Title         string
	Link          string
	Description   string
	Content       string
	Creator       string
	Date          time.Time
	BookmarkCount int
	Subject       []string
}

func newFavoriteFeed(f UnifiedFeed) FavoriteFeed {
	ret := FavoriteFeed{
		Title:       f.Title,
		Link:        f.Link,
		Description: f.Description,
		Items:       newFavoriteFeedItems(f.Items),
	}
	return ret
}

func newFavoriteFeedItems(items []UnifiedFeedItem) []FavoriteFeedItem {
	ret := []FavoriteFeedItem{}
	for _, i := range items {
		ret = append(ret, FavoriteFeedItem{
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
