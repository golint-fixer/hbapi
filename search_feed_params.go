package hbapi

import (
	"strconv"
	"time"
)

type targetType int

const (
	targetTag targetType = iota
	targetTitle
	targetText
)

func (t targetType) String() string {
	switch t {
	case targetTag:
		return "tag"
	case targetTitle:
		return "title"
	case targetText:
		return "text"
	default:
		return ""
	}
}

type sortType int

const (
	sortRecent sortType = iota
	sortPopular
)

func (t sortType) String() string {
	switch t {
	case sortRecent:
		return "recent"
	case sortPopular:
		return "popular"
	default:
		return ""
	}
}

type safeType int

const (
	safeOn safeType = iota
	safeOff
)

func (t safeType) String() string {
	switch t {
	case safeOn:
		return "on"
	case safeOff:
		return "off"
	default:
		return ""
	}
}

// SearchFeedParams represents request parameters for hatena bookmark search.
type SearchFeedParams struct {
	query     string
	target    targetType
	sort      sortType
	users     int
	safe      safeType
	dateBegin time.Time
	dateEnd   time.Time
}

// NewSearchFeedParams initialize SearchFeedParams.
func NewSearchFeedParams(query string) SearchFeedParams {
	params := SearchFeedParams{
		query:  query,
		target: targetText,
		sort:   sortRecent,
		users:  3,
		safe:   safeOn,
	}
	return params
}

// GetRequest return request url for search.
func (params SearchFeedParams) GetRequest() string {
	req := "http://b.hatena.ne.jp/search/"

	req += params.target.String()
	req += "?q=" + params.query
	req += "&sort=" + params.sort.String()
	req += "&users=" + strconv.Itoa(params.users)
	req += "&safe=" + params.safe.String()

	if !params.dateEnd.IsZero() {
		// yyyy-MM-dd
		req += "&date_end=" + params.dateEnd.Format("2006-01-02")

	}

	if !params.dateBegin.IsZero() {
		// yyyy-MM-dd
		req += "&date_begin=" + params.dateBegin.Format("2006-01-02")
	}

	req += "&mode=rss"

	return req
}

// SetQuery set query.
func (params *SearchFeedParams) SetQuery(query string) {
	if query != "" {
		params.query = query
	}
}

// SetTarget2Tag set target to tag.
func (params *SearchFeedParams) SetTarget2Tag() {
	params.target = targetTag
}

// SetTarget2Title set target to title.
func (params *SearchFeedParams) SetTarget2Title() {
	params.target = targetTitle
}

// SetTarget2Text set target text.
func (params *SearchFeedParams) SetTarget2Text() {
	params.target = targetText
}

// SetSort2Recent set sort to recent.
func (params *SearchFeedParams) SetSort2Recent() {
	params.sort = sortRecent
}

// SetSort2Popular set sort to popuser.
func (params *SearchFeedParams) SetSort2Popular() {
	params.sort = sortPopular
}

// SetUsers set users.
func (params *SearchFeedParams) SetUsers(users int) {
	if users >= 0 {
		params.users = users
	}
}

// EnableSafe set safe true.
func (params *SearchFeedParams) EnableSafe() {
	params.safe = safeOn
}

// DisableSafe set safe false.
func (params *SearchFeedParams) DisableSafe() {
	params.safe = safeOff
}

// SetDateBegin set date_begin.
func (params *SearchFeedParams) SetDateBegin(date time.Time) {
	if !date.IsZero() {
		params.dateBegin = date
	}
}

// SetDateEnd set date_end.
func (params *SearchFeedParams) SetDateEnd(date time.Time) {
	if !date.IsZero() {
		params.dateEnd = date
	}
}
