package hbapi

import (
	"strconv"
)

// HBFeedParams represents request parameters for get feed.
type HBFeedParams struct {
	User string
	Page int
	Tag  string
	Date string // yyyyMMdd
	URL  string
}

// NewHBFeedParams initialize HBFeedParams.
func NewHBFeedParams(user string) HBFeedParams {
	params := HBFeedParams{
		User: user,
		Page: 0,
		Tag:  "",
		Date: "",
		URL:  "",
	}
	return params
}

// GetRequest return request url for get feed.
func (params HBFeedParams) GetRequest() string {
	req := "http://b.hatena.ne.jp/" + params.User + "/rss"

	req += "?of=" + strconv.Itoa(20*params.Page)

	if params.Tag != "" {
		req += "&tag=" + params.Tag
	}

	if params.Date != "" {
		req += "&date=" + params.Date
	}

	if params.URL != "" {
		req += "&url=" + params.URL
	}

	return req
}
