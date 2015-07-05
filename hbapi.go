// "hbapi" is library for hatena bookmark api.
package hbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ENTRY_INFO_API_URL   = "http://b.hatena.ne.jp/entry/json/"
	ENTRY_COUNT_API_URL  = "http://api.b.st-hatena.com/entry.count"
	ENTRY_COUNTS_API_URL = "http://api.b.st-hatena.com/entry.counts"
)

// call hatena bookmark entry information api
func GetEntryInfo(target string) (HBEntryInfo, error) {
	q, err := url.Parse(target)
	if err != nil {
		return HBEntryInfo{}, err
	}

	v := url.Values{}
	v.Set("url", q.String())

	u, _ := url.Parse(ENTRY_INFO_API_URL)
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

// call hatena bookmark count api(single url)
func GetBookmarkCount(target string) (int, error) {
	q, err := url.Parse(target)
	if err != nil {
		return 0, err
	}

	v := url.Values{}
	v.Set("url", q.String())

	u, _ := url.Parse(ENTRY_COUNT_API_URL)
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

// call hatena bookmark count api(multiple url)
func GetBookmarkCounts(targets []string) (map[string]int, error) {
	v := url.Values{}
	for _, target := range targets {
		q, err := url.Parse(target)
		if err != nil {
			return map[string]int{}, err
		}
		v.Add("url", q.String())
	}

	u, _ := url.Parse(ENTRY_COUNTS_API_URL)
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
