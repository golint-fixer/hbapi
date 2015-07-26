package hbapi

import (
	"strconv"
	"time"
)

// HBFavoriteFeedParams represents request parameters for get favorite feed.
type HBFavoriteFeedParams struct {
	User   string
	Until  time.Time
	Page   int
	WithMe bool
}

// NewHBFavoriteFeedParams initialize HBFavoriteFeedParams.
func NewHBFavoriteFeedParams(user string) HBFavoriteFeedParams {
	params := HBFavoriteFeedParams{
		User:   user,
		Page:   0,
		WithMe: false,
	}
	return params
}

// GetRequest return request url for get favorite feed.
func (params HBFavoriteFeedParams) GetRequest() string {
	req := "http://b.hatena.ne.jp/" + params.User + "/favorite.rss"

	// MEMO untilとofは同時に使えない？
	if !params.Until.IsZero() {
		req += "?until=" + strconv.Itoa(int(params.Until.Unix()))
	} else {
		req += "?of=" + strconv.Itoa(25*params.Page)
	}
	if params.WithMe {
		req += "&with_me=1"
	}
	return req
}
