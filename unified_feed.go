package hbapi

import (
	"time"
)

// UnifiedFeed represents hatena bookmark unified feed.
type UnifiedFeed struct {
	Title        string
	Link         string
	Description  string
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Items        []UnifiedFeedItem
}

// UnifiedFeedItem represents hatena bookmark unified feed item.
type UnifiedFeedItem struct {
	Title         string
	Link          string
	Description   string
	Content       string
	Creator       string
	Date          time.Time
	BookmarkCount int
	Subject       []string
}
