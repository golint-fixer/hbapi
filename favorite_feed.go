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
