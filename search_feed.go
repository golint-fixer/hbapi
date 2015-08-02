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
