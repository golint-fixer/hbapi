package hbapi

import (
	"strconv"
	"time"
)

// Feed items per page.
const (
	FeedItemsPerPage = 20
)

// HBFeedParams represents request parameters for get feed.
type HBFeedParams struct {
	user string
	page int
	tag  string
	date time.Time
	url  string
}

// NewHBFeedParams initialize HBFeedParams.
func NewHBFeedParams(user string) HBFeedParams {
	params := HBFeedParams{
		user: user,
		page: 0,
		tag:  "",
		url:  "",
	}
	return params
}

// GetRequest return request url for get feed.
func (params HBFeedParams) GetRequest() string {
	req := "http://b.hatena.ne.jp/" + params.user + "/rss"

	req += "?of=" + strconv.Itoa(FeedItemsPerPage*params.page)

	if params.tag != "" {
		req += "&tag=" + params.tag
	}

	if !params.date.IsZero() {
		// date format: yyyyMMdd
		req += "&date=" + params.date.Format("20060702")
	}

	if params.url != "" {
		req += "&url=" + params.url
	}

	return req
}

// SetUser set user param.
func (params *HBFeedParams) SetUser(user string) {
	if user != "" {
		params.user = user
	}
}

// SetPage set page param.
func (params *HBFeedParams) SetPage(page int) {
	if page >= 0 {
		params.page = page
	}
}

// SetTag set tag param.
func (params *HBFeedParams) SetTag(tag string) {
	if tag != "" {
		params.tag = tag
	}
}

// SetDate set date param.
func (params *HBFeedParams) SetDate(date time.Time) {
	if !date.IsZero() {
		params.date = date
	}
}

// SetURL set url param.
func (params *HBFeedParams) SetURL(url string) {
	if url != "" {
		params.url = url
	}
}
