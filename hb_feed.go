package hbapi

import (
	"time"
)

// HBFeed represents feed.
type HBFeed struct {
	Title        string
	Link         string
	Description  string
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Items        []HBFeedItem
}

// HBFeedItem represents feed items.
type HBFeedItem struct {
	Title         string
	Link          string
	Description   string
	Content       string
	Creator       string
	Date          time.Time
	BookmarkCount int
	Subject       []string
}
