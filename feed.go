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
