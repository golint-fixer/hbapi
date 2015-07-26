package hbapi

import (
	"time"
)

// HBFavoriteFeed represents favorite feed.
type HBFavoriteFeed struct {
	Title       string
	Link        string
	Description string
	Items       []HBFavoriteFeedItem
}

// HBFavoriteFeedItem represents favorite feed items.
type HBFavoriteFeedItem struct {
	Title         string
	Link          string
	Description   string
	Content       string
	Creator       string
	Date          time.Time
	BookmarkCount int
	Subject       []string
}
