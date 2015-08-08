package hbapi

import (
	neturl "net/url"
	"strconv"
	"time"
)

// Feed items per page.
const (
	FeedItemsPerPage = 20
)

// FeedParams represents request parameters for get hatena bookmark feed.
type FeedParams struct {
	user string
	page int
	tag  string
	date time.Time
	url  string
}

// NewFeedParams initialize FeedParams.
func NewFeedParams(user string) FeedParams {
	params := FeedParams{
		user: user,
		page: 0,
		tag:  "",
		url:  "",
	}
	return params
}

// GetRequest return request url for get hatena bookmark feed.
func (params FeedParams) GetRequest() string {
	query := neturl.Values{}
	query.Set("of", strconv.Itoa(FeedItemsPerPage*params.page))

	if params.tag != "" {
		query.Set("tag", params.tag)
	}

	if !params.date.IsZero() {
		// date format: yyyyMMdd
		query.Set("date", params.date.Format("20060102"))
	}

	if params.url != "" {
		query.Set("url", params.url)
	}

	url, _ := neturl.Parse("http://b.hatena.ne.jp/" + params.user + "/rss")
	url.RawQuery = query.Encode()

	return url.String()
}

// SetUser set user param.
func (params *FeedParams) SetUser(user string) {
	if user != "" {
		params.user = user
	}
}

// SetPage set page param.
func (params *FeedParams) SetPage(page int) {
	if page >= 0 {
		params.page = page
	}
}

// SetTag set tag param.
func (params *FeedParams) SetTag(tag string) {
	if tag != "" {
		params.tag = tag
	}
}

// SetDate set date param.
func (params *FeedParams) SetDate(date time.Time) {
	if !date.IsZero() {
		params.date = date
	}
}

// SetURL set url param.
func (params *FeedParams) SetURL(url string) {
	if url != "" {
		params.url = url
	}
}
