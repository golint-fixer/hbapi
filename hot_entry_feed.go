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
