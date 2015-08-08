package hbapi

import (
	neturl "net/url"
	"strconv"
	"time"
)

// Favorite feed items per page.
const (
	FavoriteFeedItemsPerPage = 25
)

// FavoriteFeedParams represents request parameters for get hatena bookmark favorite feed.
type FavoriteFeedParams struct {
	user   string
	until  time.Time
	page   int
	withMe bool
}

// NewFavoriteFeedParams initialize FavoriteFeedParams.
func NewFavoriteFeedParams(user string) FavoriteFeedParams {
	params := FavoriteFeedParams{
		user:   user,
		page:   0,
		withMe: false,
	}
	return params
}

// GetRequest return request url for get favorite feed.
func (params FavoriteFeedParams) GetRequest() string {
	query := neturl.Values{}

	if !params.until.IsZero() {
		query.Set("until", strconv.Itoa(int(params.until.Unix())))
	} else {
		query.Set("of", strconv.Itoa(FavoriteFeedItemsPerPage*params.page))
	}

	if params.withMe {
		query.Set("with_me", "1")
	}

	url, _ := neturl.Parse("http://b.hatena.ne.jp/" + params.user + "/favorite.rss")
	url.RawQuery = query.Encode()

	return url.String()
}

// SetUser set user param.
func (params *FavoriteFeedParams) SetUser(user string) {
	if user != "" {
		params.user = user
	}
}

// SetPage set page param.
func (params *FavoriteFeedParams) SetPage(page int) {
	if page >= 0 {
		params.page = page
	}
}

// SetUntil set until param.
func (params *FavoriteFeedParams) SetUntil(until time.Time) {
	if !until.IsZero() {
		params.until = until
	}
}

// EnableWithMe set withMe param true.
func (params *FavoriteFeedParams) EnableWithMe() {
	params.withMe = true
}

// DisableWithMe set withMe param false.
func (params *FavoriteFeedParams) DisableWithMe() {
	params.withMe = false
}
