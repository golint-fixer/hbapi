// Package hbapi : library for hatena bookmark api.
package hbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// API Endpoint
const (
	EntryInfoAPIURL   = "http://b.hatena.ne.jp/entry/json/"
	EntryCountAPIURL  = "http://api.b.st-hatena.com/entry.count"
	EntryCountsAPIURL = "http://api.b.st-hatena.com/entry.counts"
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

// Search search hatena bookmark.
func Search(params SearchFeedParams) (SearchFeed, error) {
	req := params.GetRequest()

	p := newParser()
	f, err := p.Parse(req)
	if err != nil {
		return SearchFeed{}, err
	}

	return newSearchFeed(f), nil
}
