package hbapi

import (
	"testing"
	"time"
)

func TestSearchFeedParamsGetRequest(t *testing.T) {
	params := NewSearchFeedParams("golang")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithQuery(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetQuery("ruby")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=ruby&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithTargetTag(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetTarget2Tag()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/tag?q=golang&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithTargetTitle(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetTarget2Title()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/title?q=golang&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithTargetText(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetTarget2Text()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithSortRecent(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetSort2Recent()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithSortPopular(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetSort2Popular()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=popular&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithSortUsers(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetUsers(100)
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=100&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithEnableSafe(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.EnableSafe()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithDisableSafe(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.DisableSafe()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=off&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithDateBegin(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetDateBegin(time.Date(2015, time.August, 1, 12, 34, 56, 0, time.UTC))
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=on&date_begin=2015-08-01&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithDateEnd(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetDateEnd(time.Date(2015, time.August, 1, 12, 34, 56, 0, time.UTC))
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=on&date_end=2015-08-01&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
