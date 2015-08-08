package hbapi

import (
	"testing"
	"time"
)

func TestSearchFeedParamsGetRequest(t *testing.T) {
	params := NewSearchFeedParams("golang")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithQuery(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetQuery("ruby")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=ruby&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithTargetTag(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetTarget2Tag()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/tag?mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithTargetTitle(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetTarget2Title()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/title?mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithTargetText(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetTarget2Text()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithSortRecent(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetSort2Recent()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithSortPopular(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetSort2Popular()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=golang&safe=on&sort=popular&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithSortUsers(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetUsers(100)
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=golang&safe=on&sort=recent&users=100"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithEnableSafe(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.EnableSafe()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithDisableSafe(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.DisableSafe()
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?mode=rss&q=golang&safe=off&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithDateBegin(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetDateBegin(time.Date(2015, time.August, 1, 12, 34, 56, 0, time.UTC))
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?date_begin=2015-08-01&mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestSearchFeedParamsGetRequestWithDateEnd(t *testing.T) {
	params := NewSearchFeedParams("golang")
	params.SetDateEnd(time.Date(2015, time.August, 1, 12, 34, 56, 0, time.UTC))
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?date_end=2015-08-01&mode=rss&q=golang&safe=on&sort=recent&users=3"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
