// Package hbapi : library for hatena bookmark api.
package hbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
)

// API Endpoint
const (
	EntryInfoAPIURL   = "http://b.hatena.ne.jp/entry/json/"
	EntryCountAPIURL  = "http://api.b.st-hatena.com/entry.count"
	EntryCountsAPIURL = "http://api.b.st-hatena.com/entry.counts"
	HotEntryFeedURL   = "https://feeds.feedburner.com/hatena/b/hotentry"
)

// XML Namespace
const (
	XMLNSContent    = "http://purl.org/rss/1.0/modules/content/"
	XMLNSOpenSearch = "http://a9.com/-/spec/opensearchrss/1.0/"
	XMLNSDC         = "http://purl.org/dc/elements/1.1/"
	XMLNSHatena     = "http://www.hatena.ne.jp/info/xmlns#"
)

// XML Node
const (
	XMLNodeStartIndex    = "startIndex"
	XMLNodeItemsPerPage  = "itemsPerPage"
	XMLNodeTotalResults  = "totalResults"
	XMLNodeEncoded       = "encoded"
	XMLNodeCreator       = "creator"
	XMLNodeDate          = "date"
	XMLNodeBookmarkCount = "bookmarkcount"
	XMLNodeSubject       = "subject"
)

// HTTP timeout threashold
const (
	Timeout = 10
)

// GetEntryInfo call hatena bookmark entry information api.
func GetEntryInfo(target string) (EntryInfo, error) {
	q, err := url.Parse(target)
	if err != nil {
		return EntryInfo{}, err
	}

	v := url.Values{}
	v.Set("url", q.String())

	u, _ := url.Parse(EntryInfoAPIURL)
	u.RawQuery = v.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return EntryInfo{}, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	info := EntryInfo{}
	json.Unmarshal(body, &info)
	return info, nil
}

// GetBookmarkCount call hatena bookmark count api.
func GetBookmarkCount(target string) (int, error) {
	q, err := url.Parse(target)
	if err != nil {
		return 0, err
	}

	v := url.Values{}
	v.Set("url", q.String())

	u, _ := url.Parse(EntryCountAPIURL)
	u.RawQuery = v.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return 0, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	count, _ := strconv.Atoi(string(body))
	return count, nil
}

// GetBookmarkCounts call hatena bookmark count api.
func GetBookmarkCounts(targets []string) (map[string]int, error) {
	v := url.Values{}
	for _, target := range targets {
		q, err := url.Parse(target)
		if err != nil {
			return map[string]int{}, err
		}
		v.Add("url", q.String())
	}

	u, _ := url.Parse(EntryCountsAPIURL)
	u.RawQuery = v.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return map[string]int{}, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	counts := map[string]int{}
	json.Unmarshal(body, &counts)
	return counts, nil
}

// GetFeed get hatena bookmark feed.
func GetFeed(params FeedParams) (Feed, error) {
	feed := rss.New(Timeout, true, nil, nil)

	req := params.GetRequest()
	if err := feed.Fetch(req, nil); err != nil {
		return Feed{}, err
	}

	channel := feed.Channels[0]
	opensearch := channel.Extensions[XMLNSOpenSearch]

	hbf := Feed{}
	hbf.Title = channel.Title
	hbf.Link = channel.Links[0].Href
	hbf.Description = channel.Description
	if opensearch[XMLNodeStartIndex] != nil {
		startIndex, _ := strconv.Atoi(opensearch[XMLNodeStartIndex][0].Value)
		hbf.StartIndex = startIndex
	}
	if opensearch[XMLNodeItemsPerPage] != nil {
		itemsPerPage, _ := strconv.Atoi(opensearch[XMLNodeItemsPerPage][0].Value)
		hbf.ItemsPerPage = itemsPerPage
	}
	if opensearch[XMLNodeTotalResults] != nil {
		totalResults, _ := strconv.Atoi(opensearch[XMLNodeTotalResults][0].Value)
		hbf.TotalResults = totalResults
	}

	items := []FeedItem{}
	for _, item := range channel.Items {
		content := item.Extensions[XMLNSContent]
		dc := item.Extensions[XMLNSDC]
		hatena := item.Extensions[XMLNSHatena]

		i := FeedItem{}
		i.Title = item.Title
		i.Link = item.Links[0].Href
		i.Description = item.Description
		if content[XMLNodeEncoded] != nil {
			i.Content = content[XMLNodeEncoded][0].Value
		}
		if dc[XMLNodeCreator] != nil {
			i.Creator = dc[XMLNodeCreator][0].Value
		}
		if dc[XMLNodeDate] != nil {
			date, _ := time.Parse(time.RFC3339, dc[XMLNodeDate][0].Value)
			i.Date = date
		}
		if hatena[XMLNodeBookmarkCount] != nil {
			bookmarkCount, _ := strconv.Atoi(hatena[XMLNodeBookmarkCount][0].Value)
			i.BookmarkCount = bookmarkCount
		}
		if dc[XMLNodeSubject] != nil {
			for _, subject := range dc[XMLNodeSubject] {
				i.Subject = append(i.Subject, subject.Value)
			}
		}
		items = append(items, i)
	}
	hbf.Items = items

	return hbf, nil
}

// GetFavoriteFeed get hatena bookmark favorite feed.
func GetFavoriteFeed(params FavoriteFeedParams) (FavoriteFeed, error) {
	feed := rss.New(Timeout, true, nil, nil)

	req := params.GetRequest()
	if err := feed.Fetch(req, nil); err != nil {
		return FavoriteFeed{}, err
	}

	channel := feed.Channels[0]

	hbf := FavoriteFeed{}
	hbf.Title = channel.Title
	hbf.Link = channel.Links[0].Href
	hbf.Description = channel.Description

	items := []FavoriteFeedItem{}
	for _, item := range channel.Items {
		content := item.Extensions[XMLNSContent]
		dc := item.Extensions[XMLNSDC]
		hatena := item.Extensions[XMLNSHatena]

		i := FavoriteFeedItem{}
		i.Title = item.Title
		i.Link = item.Links[0].Href
		i.Description = item.Description
		if content[XMLNodeEncoded] != nil {
			i.Content = content[XMLNodeEncoded][0].Value
		}
		if dc[XMLNodeCreator] != nil {
			i.Creator = dc[XMLNodeCreator][0].Value
		}
		if dc[XMLNodeDate] != nil {
			date, _ := time.Parse(time.RFC3339, dc[XMLNodeDate][0].Value)
			i.Date = date
		}
		if hatena[XMLNodeBookmarkCount] != nil {
			bookmarkCount, _ := strconv.Atoi(hatena[XMLNodeBookmarkCount][0].Value)
			i.BookmarkCount = bookmarkCount
		}
		if dc[XMLNodeSubject] != nil {
			for _, subject := range dc[XMLNodeSubject] {
				i.Subject = append(i.Subject, subject.Value)
			}
		}
		items = append(items, i)
	}
	hbf.Items = items

	return hbf, nil
}

// GetHotEntryFeed get hatena bookmark hot entry feed.
func GetHotEntryFeed() (HotEntryFeed, error) {
	feed := rss.New(Timeout, true, nil, nil)

	if err := feed.Fetch(HotEntryFeedURL, nil); err != nil {
		return HotEntryFeed{}, err
	}

	channel := feed.Channels[0]

	hbf := HotEntryFeed{}
	hbf.Title = channel.Title
	hbf.Link = channel.Links[0].Href
	hbf.Description = channel.Description

	items := []HotEntryFeedItem{}
	for _, item := range channel.Items {
		content := item.Extensions[XMLNSContent]
		dc := item.Extensions[XMLNSDC]
		hatena := item.Extensions[XMLNSHatena]

		i := HotEntryFeedItem{}
		i.Title = item.Title
		i.Link = item.Links[0].Href
		i.Description = item.Description
		if content[XMLNodeEncoded] != nil {
			i.Content = content[XMLNodeEncoded][0].Value
		}
		if dc[XMLNodeCreator] != nil {
			i.Creator = dc[XMLNodeCreator][0].Value
		}
		if dc[XMLNodeDate] != nil {
			date, _ := time.Parse(time.RFC3339, dc[XMLNodeDate][0].Value)
			i.Date = date
		}
		if hatena[XMLNodeBookmarkCount] != nil {
			bookmarkCount, _ := strconv.Atoi(hatena[XMLNodeBookmarkCount][0].Value)
			i.BookmarkCount = bookmarkCount
		}
		if dc[XMLNodeSubject] != nil {
			for _, subject := range dc[XMLNodeSubject] {
				i.Subject = append(i.Subject, subject.Value)
			}
		}
		items = append(items, i)
	}
	hbf.Items = items

	return hbf, nil
}

// Search search hatena bookmark.
func Search(params SearchFeedParams) (SearchFeed, error) {
	feed := rss.New(Timeout, true, nil, nil)

	if err := feed.Fetch(params.GetRequest(), nil); err != nil {
		return SearchFeed{}, err
	}

	channel := feed.Channels[0]

	hbf := SearchFeed{}
	hbf.Title = channel.Title
	hbf.Link = channel.Links[0].Href
	hbf.Description = channel.Description

	items := []SearchFeedItem{}
	for _, item := range channel.Items {
		content := item.Extensions[XMLNSContent]
		dc := item.Extensions[XMLNSDC]
		hatena := item.Extensions[XMLNSHatena]

		i := SearchFeedItem{}
		i.Title = item.Title
		i.Link = item.Links[0].Href
		i.Description = item.Description
		if content[XMLNodeEncoded] != nil {
			i.Content = content[XMLNodeEncoded][0].Value
		}
		if dc[XMLNodeCreator] != nil {
			i.Creator = dc[XMLNodeCreator][0].Value
		}
		if dc[XMLNodeDate] != nil {
			date, _ := time.Parse(time.RFC3339, dc[XMLNodeDate][0].Value)
			i.Date = date
		}
		if hatena[XMLNodeBookmarkCount] != nil {
			bookmarkCount, _ := strconv.Atoi(hatena[XMLNodeBookmarkCount][0].Value)
			i.BookmarkCount = bookmarkCount
		}
		if dc[XMLNodeSubject] != nil {
			for _, subject := range dc[XMLNodeSubject] {
				i.Subject = append(i.Subject, subject.Value)
			}
		}
		items = append(items, i)
	}
	hbf.Items = items

	return hbf, nil
}
