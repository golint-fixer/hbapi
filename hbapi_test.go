package hbapi

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"strconv"
	"testing"
)

func TestGetEntryInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := `
{
	"title":"\u306f\u3066\u306a",
	"count":"5523",
	"url":"http://www.hatena.ne.jp/",
	"entry_url":"http://b.hatena.ne.jp/entry/www.hatena.ne.jp/",
	"eid":"370",
	"screenshot":"http://screenshot.hatena.ne.jp/images/200x150/c/2/6/b/d/f015f87fcf44a513f4744d853fe22504bc3.jpg",
	"bookmarks":[
		{
			"user":"user01",
		  "tags":["tag01", "tag02"],
			"timestamp":"2015/07/01 00:00:00",
			"comment":"comment01"
		},
		{
			"user":"user02",
			"tags":["tag03", "tag04"],
			"timestamp":"2015/07/02 00:00:00",
			"comment":"comment02"
		}
	],
	"related":[
		{
			"title":"title01",
			"count":"1",
			"url":"http://url01.example.com",
			"entry_url":"http://entry_url01.example.com",
			"eid":"1"
		},
		{
			"title":"title02",
			"count":"2",
			"url":"http://url02.example.com",
			"entry_url":"http://entry_url02.example.com",
			"eid":"2"
		}
	]
}
`
	httpmock.RegisterResponder(
		"GET",
		"http://b.hatena.ne.jp/entry/json/?url=http%3A%2F%2Fdeveloper.hatena.ne.jp",
		httpmock.NewStringResponder(200, res))

	input := "http://developer.hatena.ne.jp"
	expected := HBEntryInfo{}
	json.Unmarshal([]byte(res), &expected)

	info, err := GetEntryInfo(input)
	if err != nil {
		t.Errorf("fail mock: %s\n", input)
	}

	// Title
	if info.Title != expected.Title {
		t.Errorf("expected title %s, but got %s\n", expected.Title, info.Title)
	}

	// Count
	if info.Count != expected.Count {
		t.Errorf("expected count %s, but got %s\n", expected.Count, info.Count)
	}

	// Url
	if info.Url != expected.Url {
		t.Errorf("expected url %s, but got %s\n", expected.Url, info.Url)
	}

	// EntryUrl
	if info.EntryUrl != expected.EntryUrl {
		t.Errorf("expected entry_url %s, but got %s\n", expected.EntryUrl, info.EntryUrl)
	}

	// Eid
	if info.Eid != expected.Eid {
		t.Errorf("expected eid %s, but got %s\n", expected.Eid, info.Eid)
	}

	// Screenshot
	if info.Screenshot != expected.Screenshot {
		t.Errorf("expected screenshot %s, but got %s\n", expected.Screenshot, info.Screenshot)
	}

	// Bookmarks
	for i, _ := range expected.Bookmarks {
		// User
		if info.Bookmarks[i].User != expected.Bookmarks[i].User {
			t.Errorf("expected bookmarks user %s, but got %s\n", expected.Bookmarks[i].User, info.Bookmarks[i].User)
		}

		// Tags
		for j, _ := range expected.Bookmarks[i].Tags {
			if info.Bookmarks[i].Tags[j] != expected.Bookmarks[i].Tags[j] {
				t.Errorf("expected bookmarks tags %s, but got %s\n", expected.Bookmarks[i].Tags[j], info.Bookmarks[i].Tags[j])
			}
		}

		// Timestamp
		if info.Bookmarks[i].Timestamp != expected.Bookmarks[i].Timestamp {
			t.Errorf("expected bookmarks timestamp %s, but got %s\n", expected.Bookmarks[i].Timestamp, info.Bookmarks[i].Timestamp)
		}

		// Comment
		if info.Bookmarks[i].Comment != expected.Bookmarks[i].Comment {
			t.Errorf("expected bookmarks comment %s, but got %s\n", expected.Bookmarks[i].Comment, info.Bookmarks[i].Comment)
		}
	}

	// Related
	for i, _ := range expected.Related {
		// Title
		if info.Related[i].Title != expected.Related[i].Title {
			t.Errorf("expected related title %s, but got %s\n", expected.Related[i].Title, info.Related[i].Title)
		}

		// Count
		if info.Related[i].Count != expected.Related[i].Count {
			t.Errorf("expected related count %s, but got %s\n", expected.Related[i].Count, info.Related[i].Count)
		}

		// Url
		if info.Related[i].Url != expected.Related[i].Url {
			t.Errorf("expected related url %s, but got %s\n", expected.Related[i].Url, info.Related[i].Url)
		}

		// EntryUrl
		if info.Related[i].EntryUrl != expected.Related[i].EntryUrl {
			t.Errorf("expected related entry_url %s, but got %s\n", expected.Related[i].EntryUrl, info.Related[i].EntryUrl)
		}

		// Eid
		if info.Related[i].Eid != expected.Related[i].Eid {
			t.Errorf("expected related eid %s, but got %s\n", expected.Related[i].Eid, info.Related[i].Eid)
		}
	}
}

func TestGetBookmarkCount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := "999"
	httpmock.RegisterResponder(
		"GET",
		"http://api.b.st-hatena.com/entry.count?url=http%3A%2F%2Fdeveloper.hatena.ne.jp",
		httpmock.NewStringResponder(200, res))

	input := "http://developer.hatena.ne.jp"
	expected, _ := strconv.Atoi(res)

	count, err := GetBookmarkCount(input)
	if err != nil {
		t.Errorf("fail mock: %s\n", input)
	}

	if count != expected {
		t.Errorf("expected count %d, but got %d\n", expected, count)
	}
}

func TestGetBookmarkCounts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := `
{
	"https://github.com":111,
	"https://bitbucket.org":222,
	"http://stackoverflow.com":333
}
`
	httpmock.RegisterResponder(
		"GET",
		"http://api.b.st-hatena.com/entry.counts?url=https%3A%2F%2Fgithub.com&url=https%3A%2F%2Fbitbucket.org&url=http%3A%2F%2Fstackoverflow.com",
		httpmock.NewStringResponder(200, res))

	input := []string{
		"https://github.com",
		"https://bitbucket.org",
		"http://stackoverflow.com",
	}
	expected := map[string]int{}
	json.Unmarshal([]byte(res), &expected)

	counts, err := GetBookmarkCounts(input)
	if err != nil {
		t.Errorf("fail mock: %#v\n", input)
	}

	for url, count := range counts {
		if count != expected[url] {
			t.Errorf("expected count %d, but got %d\n", expected[url], count)
		}
	}
}
