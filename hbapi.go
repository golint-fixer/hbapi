// Package hbapi : library for hatena bookmark api.
package hbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"strconv"
)

// API Endpoint
const (
	EntryInfoAPIURL   = "http://b.hatena.ne.jp/entry/json/"
	EntryCountAPIURL  = "http://api.b.st-hatena.com/entry.count"
	EntryCountsAPIURL = "http://api.b.st-hatena.com/entry.counts"
)

// GetEntryInfo call hatena bookmark entry information api.
func GetEntryInfo(url string) (EntryInfo, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return EntryInfo{}, err
	}

	query := neturl.Values{}
	query.Set("url", u.String())

	req, _ := neturl.Parse(EntryInfoAPIURL)
	req.RawQuery = query.Encode()

	res, err := http.Get(req.String())
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
func GetBookmarkCount(url string) (int, error) {
	u, err := neturl.Parse(url)
	if err != nil {
		return 0, err
	}

	query := neturl.Values{}
	query.Set("url", u.String())

	req, _ := neturl.Parse(EntryCountAPIURL)
	req.RawQuery = query.Encode()

	res, err := http.Get(req.String())
	if err != nil {
		return 0, err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	count, _ := strconv.Atoi(string(body))
	return count, nil
}

// GetBookmarkCounts call hatena bookmark count api.
func GetBookmarkCounts(urls []string) (map[string]int, error) {
	query := neturl.Values{}
	for _, url := range urls {
		u, err := neturl.Parse(url)
		if err != nil {
			return map[string]int{}, err
		}
		query.Add("url", u.String())
	}

	req, _ := neturl.Parse(EntryCountsAPIURL)
	req.RawQuery = query.Encode()

	res, err := http.Get(req.String())
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
	req := params.GetRequest()

	p := newParser()
	f, err := p.Parse(req)
	if err != nil {
		return Feed{}, err
	}

	return newFeed(f), nil
}

// GetFavoriteFeed get hatena bookmark favorite feed.
func GetFavoriteFeed(params FavoriteFeedParams) (FavoriteFeed, error) {
	req := params.GetRequest()

	p := newParser()
	f, err := p.Parse(req)
	if err != nil {
		return FavoriteFeed{}, err
	}

	return newFavoriteFeed(f), nil
}

// GetHotEntryFeed get hatena bookmark hot entry feed.
func GetHotEntryFeed(params HotEntryFeedParams) (HotEntryFeed, error) {
	req := params.GetRequest()

	p := newParser()
	f, err := p.Parse(req)
	if err != nil {
		return HotEntryFeed{}, err
	}

	return newHotEntryFeed(f), nil
}

// GetSearchFeed search hatena bookmark.
func GetSearchFeed(params SearchFeedParams) (SearchFeed, error) {
	req := params.GetRequest()

	p := newParser()
	f, err := p.Parse(req)
	if err != nil {
		return SearchFeed{}, err
	}

	return newSearchFeed(f), nil
}
