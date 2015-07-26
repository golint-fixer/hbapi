package hbapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestGetEntryInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := `
{
	"title":"\u306f\u3066\u306a",
	"count":5523,
	"url":"http://www.hatena.ne.jp/",
	"entry_url":"http://b.hatena.ne.jp/entry/www.hatena.ne.jp/",
	"eid":370,
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
			"count":1,
			"url":"http://url01.example.com",
			"entry_url":"http://entry_url01.example.com",
			"eid":1
		},
		{
			"title":"title02",
			"count":2,
			"url":"http://url02.example.com",
			"entry_url":"http://entry_url02.example.com",
			"eid":2
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
		t.Errorf("expected count %d, but got %d\n", expected.Count, info.Count)
	}

	// URL
	if info.URL != expected.URL {
		t.Errorf("expected url %s, but got %s\n", expected.URL, info.URL)
	}

	// EntryURL
	if info.EntryURL != expected.EntryURL {
		t.Errorf("expected entry_url %s, but got %s\n", expected.EntryURL, info.EntryURL)
	}

	// Eid
	if info.Eid != expected.Eid {
		t.Errorf("expected eid %d, but got %d\n", expected.Eid, info.Eid)
	}

	// Screenshot
	if info.Screenshot != expected.Screenshot {
		t.Errorf("expected screenshot %s, but got %s\n", expected.Screenshot, info.Screenshot)
	}

	// Bookmarks
	for i := range expected.Bookmarks {
		// User
		if info.Bookmarks[i].User != expected.Bookmarks[i].User {
			t.Errorf("expected bookmarks user %s, but got %s\n", expected.Bookmarks[i].User, info.Bookmarks[i].User)
		}

		// Tags
		for j := range expected.Bookmarks[i].Tags {
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
	for i := range expected.Related {
		// Title
		if info.Related[i].Title != expected.Related[i].Title {
			t.Errorf("expected related title %s, but got %s\n", expected.Related[i].Title, info.Related[i].Title)
		}

		// Count
		if info.Related[i].Count != expected.Related[i].Count {
			t.Errorf("expected related count %d, but got %d\n", expected.Related[i].Count, info.Related[i].Count)
		}

		// URL
		if info.Related[i].URL != expected.Related[i].URL {
			t.Errorf("expected related url %s, but got %s\n", expected.Related[i].URL, info.Related[i].URL)
		}

		// EntryURL
		if info.Related[i].EntryURL != expected.Related[i].EntryURL {
			t.Errorf("expected related entry_url %s, but got %s\n", expected.Related[i].EntryURL, info.Related[i].EntryURL)
		}

		// Eid
		if info.Related[i].Eid != expected.Related[i].Eid {
			t.Errorf("expected related eid %d, but got %d\n", expected.Related[i].Eid, info.Related[i].Eid)
		}
	}
}

func TestGetEntryInfoError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"http://b.hatena.ne.jp/entry/json/?url=http%3A%2F%2Fdeveloper.hatena.ne.jp",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("internal server error")
		},
	)

	input := "http://developer.hatena.ne.jp"
	expected := HBEntryInfo{}

	info, err := GetEntryInfo(input)
	if err == nil {
		t.Errorf("fail mock: %s\n", input)
	}

	if !reflect.DeepEqual(info, expected) {
		t.Errorf("expected %#v, but got %#v\n", expected, info)
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

func TestGetBookmarkCountError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := "0"
	httpmock.RegisterResponder(
		"GET",
		"http://api.b.st-hatena.com/entry.count?url=http%3A%2F%2Fdeveloper.hatena.ne.jp",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("internal server error")
		},
	)

	input := "http://developer.hatena.ne.jp"
	expected, err := strconv.Atoi(res)

	count, err := GetBookmarkCount(input)
	if err == nil {
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

func TestGetBookmarkCountsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"http://api.b.st-hatena.com/entry.counts?url=https%3A%2F%2Fgithub.com&url=https%3A%2F%2Fbitbucket.org&url=http%3A%2F%2Fstackoverflow.com",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("internal server error")
		},
	)

	input := []string{
		"https://github.com",
		"https://bitbucket.org",
		"http://stackoverflow.com",
	}
	expected := map[string]int{}

	counts, err := GetBookmarkCounts(input)
	if err == nil {
		t.Errorf("fail mock: %#v\n", input)
	}

	if !reflect.DeepEqual(counts, expected) {
		t.Errorf("expected counts %#v, but got %#v\n", expected, counts)
	}
}

func TestGetFeed(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	res := `
<?xml version="1.0" encoding="UTF-8"?>
<rdf:RDF
 xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
 xmlns="http://purl.org/rss/1.0/"
 xmlns:content="http://purl.org/rss/1.0/modules/content/"
 xmlns:taxo="http://purl.org/rss/1.0/modules/taxonomy/"
 xmlns:opensearch="http://a9.com/-/spec/opensearchrss/1.0/"
 xmlns:dc="http://purl.org/dc/elements/1.1/"
 xmlns:hatena="http://www.hatena.ne.jp/info/xmlns#"
 xmlns:media="http://search.yahoo.com/mrss"
>
  <channel rdf:about="http://b.hatena.ne.jp/yukihir0/">
    <title>title</title>
    <link>link</link>
    <description>description</description>
    <opensearch:startIndex>0</opensearch:startIndex>
    <opensearch:itemsPerPage>20</opensearch:itemsPerPage>
    <opensearch:totalResults>12345</opensearch:totalResults>
    <items>
      <rdf:Seq>
        <rdf:li rdf:resource="resource01" />
      </rdf:Seq>
    </items>
  </channel>
  <item rdf:about="about01">
    <title>title01</title>
    <link>link01</link>
    <description>description01</description>
    <content:encoded>content01</content:encoded>
    <dc:creator>creator01</dc:creator>
    <dc:date>2015-07-01T01:23:45+09:00</dc:date>
		<dc:subject>subject01</dc:subject>
    <taxo:topics>
      <rdf:Bag>
        <rdf:li resource="/search/tag?q=subject01" />
      </rdf:Bag>
    </taxo:topics>
    <hatena:bookmarkcount>111</hatena:bookmarkcount>
  </item>
	</rdf:RDF>
	`

	httpmock.RegisterResponder(
		"GET",
		"http://b.hatena.ne.jp/yukihir0/rss",
		httpmock.NewStringResponder(200, res))

	input := "yukihir0"
	params := NewHBFeedParams(input)
	expected := HBFeed{}
	expected.Title = "title"
	expected.Link = "link"
	expected.Description = "description"
	expected.StartIndex = 0
	expected.ItemsPerPage = 20
	expected.TotalResults = 12345
	item := HBFeedItem{}
	item.Title = "title01"
	item.Link = "link01"
	item.Description = "description01"
	item.Content = "content01"
	item.Creator = "creator01"
	date, _ := time.Parse(time.RFC3339, "2015-07-01T01:23:45+09:00")
	item.Date = date
	item.BookmarkCount = 111
	item.Subject = []string{"subject01"}
	expected.Items = append(expected.Items, item)

	feed, err := GetFeed(params)
	if err != nil {
		t.Errorf("fail mock: %#v\n", input)
	}

	if !reflect.DeepEqual(feed, expected) {
		t.Errorf("expected feed %#v, but got %#v\n", expected, feed)
	}
}

func TestGetFeedError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"http://b.hatena.ne.jp/yukihir0/rss",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("internal server error")
		},
	)

	input := "yukihir0"
	params := NewHBFeedParams(input)
	expected := HBFeed{}

	feed, err := GetFeed(params)
	if err == nil {
		t.Errorf("fail mock: %#v\n", input)
	}

	if !reflect.DeepEqual(feed, expected) {
		t.Errorf("expected feed %#v, but got %#v\n", expected, feed)
	}
}
