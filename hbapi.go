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
)

// GetEntryInfo call hatena bookmark entry information api.
func GetEntryInfo(target string) (HBEntryInfo, error) {
	q, err := url.Parse(target)
	if err != nil {
		return HBEntryInfo{}, err
	}

	v := url.Values{}
	v.Set("url", q.String())

	u, _ := url.Parse(EntryInfoAPIURL)
	u.RawQuery = v.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		return HBEntryInfo{}, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	info := HBEntryInfo{}
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
func GetFeed(params HBFeedParams) (HBFeed, error) {
	timeout := 5
	feed := rss.New(timeout, true, nil, nil)

	req := params.GetRequest()
	if err := feed.Fetch(req, nil); err != nil {
		return HBFeed{}, err
	}

	channel := feed.Channels[0]
	opensearch := channel.Extensions["http://a9.com/-/spec/opensearchrss/1.0/"]

	hbf := HBFeed{}
	hbf.Title = channel.Title
	hbf.Link = channel.Links[0].Href
	hbf.Description = channel.Description
	if opensearch["startIndex"] != nil {
		startIndex, _ := strconv.Atoi(opensearch["startIndex"][0].Value)
		hbf.StartIndex = startIndex
	}
	if opensearch["itemsPerPage"] != nil {
		itemsPerPage, _ := strconv.Atoi(opensearch["itemsPerPage"][0].Value)
		hbf.ItemsPerPage = itemsPerPage
	}
	if opensearch["totalResults"] != nil {
		totalResults, _ := strconv.Atoi(opensearch["totalResults"][0].Value)
		hbf.TotalResults = totalResults
	}

	items := []HBFeedItem{}
	for _, item := range channel.Items {
		content := item.Extensions["http://purl.org/rss/1.0/modules/content/"]
		dc := item.Extensions["http://purl.org/dc/elements/1.1/"]
		hatena := item.Extensions["http://www.hatena.ne.jp/info/xmlns#"]

		i := HBFeedItem{}
		i.Title = item.Title
		i.Link = item.Links[0].Href
		i.Description = item.Description
		if content["encoded"] != nil {
			i.Content = content["encoded"][0].Value
		}
		if dc["creator"] != nil {
			i.Creator = dc["creator"][0].Value
		}
		if dc["date"] != nil {
			date, _ := time.Parse(time.RFC3339, dc["date"][0].Value)
			i.Date = date
		}
		if hatena["bookmarkcount"] != nil {
			bookmarkCount, _ := strconv.Atoi(hatena["bookmarkcount"][0].Value)
			i.BookmarkCount = bookmarkCount
		}
		if dc["subject"] != nil {
			for _, subject := range dc["subject"] {
				i.Subject = append(i.Subject, subject.Value)
			}
		}
		items = append(items, i)
	}
	hbf.Items = items

	return hbf, nil
}

// GetFavoriteFeed get hatena bookmark favorite feed.
func GetFavoriteFeed(params HBFavoriteFeedParams) (HBFavoriteFeed, error) {
	timeout := 5
	feed := rss.New(timeout, true, nil, nil)

	req := params.GetRequest()
	if err := feed.Fetch(req, nil); err != nil {
		return HBFavoriteFeed{}, err
	}

	channel := feed.Channels[0]

	hbf := HBFavoriteFeed{}
	hbf.Title = channel.Title
	hbf.Link = channel.Links[0].Href
	hbf.Description = channel.Description

	items := []HBFavoriteFeedItem{}
	for _, item := range channel.Items {
		content := item.Extensions["http://purl.org/rss/1.0/modules/content/"]
		dc := item.Extensions["http://purl.org/dc/elements/1.1/"]
		hatena := item.Extensions["http://www.hatena.ne.jp/info/xmlns#"]

		i := HBFavoriteFeedItem{}
		i.Title = item.Title
		i.Link = item.Links[0].Href
		i.Description = item.Description
		if content["encoded"] != nil {
			i.Content = content["encoded"][0].Value
		}
		if dc["creator"] != nil {
			i.Creator = dc["creator"][0].Value
		}
		if dc["date"] != nil {
			date, _ := time.Parse(time.RFC3339, dc["date"][0].Value)
			i.Date = date
		}
		if hatena["bookmarkcount"] != nil {
			bookmarkCount, _ := strconv.Atoi(hatena["bookmarkcount"][0].Value)
			i.BookmarkCount = bookmarkCount
		}
		if dc["subject"] != nil {
			for _, subject := range dc["subject"] {
				i.Subject = append(i.Subject, subject.Value)
			}
		}
		items = append(items, i)
	}
	hbf.Items = items

	return hbf, nil
}
