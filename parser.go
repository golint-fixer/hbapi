package hbapi

import (
	"strconv"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
)

// XML Namespace
const (
	xmlNSContent    = "http://purl.org/rss/1.0/modules/content/"
	xmlNSOpenSearch = "http://a9.com/-/spec/opensearchrss/1.0/"
	xmlNSDC         = "http://purl.org/dc/elements/1.1/"
	xmlNSHatena     = "http://www.hatena.ne.jp/info/xmlns#"
)

// XML Node
const (
	xmlNodeStartIndex    = "startIndex"
	xmlNodeItemsPerPage  = "itemsPerPage"
	xmlNodeTotalResults  = "totalResults"
	xmlNodeEncoded       = "encoded"
	xmlNodeCreator       = "creator"
	xmlNodeDate          = "date"
	xmlNodeBookmarkCount = "bookmarkcount"
	xmlNodeSubject       = "subject"
)

// HTTP timeout
const (
	timeout = 10
)

type parser struct{}

func newParser() parser {
	return parser{}
}

// Parse feed.
func (p parser) Parse(req string) (UnifiedFeed, error) {
	feed := rss.New(timeout, true, nil, nil)
	if err := feed.Fetch(req, nil); err != nil {
		return UnifiedFeed{}, err
	}

	channel := feed.Channels[0]
	opensearch := channel.Extensions[xmlNSOpenSearch]

	f := UnifiedFeed{}
	f.Title = channel.Title
	f.Link = channel.Links[0].Href
	f.Description = channel.Description
	if opensearch[xmlNodeStartIndex] != nil {
		startIndex, _ := strconv.Atoi(opensearch[xmlNodeStartIndex][0].Value)
		f.StartIndex = startIndex
	}
	if opensearch[xmlNodeItemsPerPage] != nil {
		itemsPerPage, _ := strconv.Atoi(opensearch[xmlNodeItemsPerPage][0].Value)
		f.ItemsPerPage = itemsPerPage
	}
	if opensearch[xmlNodeTotalResults] != nil {
		totalResults, _ := strconv.Atoi(opensearch[xmlNodeTotalResults][0].Value)
		f.TotalResults = totalResults
	}

	items := []UnifiedFeedItem{}
	for _, item := range channel.Items {
		content := item.Extensions[xmlNSContent]
		dc := item.Extensions[xmlNSDC]
		hatena := item.Extensions[xmlNSHatena]

		i := UnifiedFeedItem{}
		i.Title = item.Title
		i.Link = item.Links[0].Href
		i.Description = item.Description
		if content[xmlNodeEncoded] != nil {
			i.Content = content[xmlNodeEncoded][0].Value
		}
		if dc[xmlNodeCreator] != nil {
			i.Creator = dc[xmlNodeCreator][0].Value
		}
		if dc[xmlNodeDate] != nil {
			date, _ := time.Parse(time.RFC3339, dc[xmlNodeDate][0].Value)
			i.Date = date
		}
		if hatena[xmlNodeBookmarkCount] != nil {
			bookmarkCount, _ := strconv.Atoi(hatena[xmlNodeBookmarkCount][0].Value)
			i.BookmarkCount = bookmarkCount
		}
		if dc[xmlNodeSubject] != nil {
			for _, subject := range dc[xmlNodeSubject] {
				i.Subject = append(i.Subject, subject.Value)
			}
		}
		items = append(items, i)
	}
	f.Items = items

	return f, nil

}
