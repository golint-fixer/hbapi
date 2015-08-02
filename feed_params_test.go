package hbapi

import (
	"testing"
	"time"
)

func TestFeedParamsGetRequest(t *testing.T) {
	params := NewFeedParams("yukihir0")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/rss?of=0"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFeedParamsGetRequestWithUser(t *testing.T) {
	params := NewFeedParams("yukihir0")
	params.SetUser("sample")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/sample/rss?of=0"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFeedParamsGetRequestWithPage(t *testing.T) {
	params := NewFeedParams("yukihir0")
	params.SetPage(1)
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/rss?of=20"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFeedParamsGetRequestWithTag(t *testing.T) {
	params := NewFeedParams("yukihir0")
	params.SetTag("golang")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/rss?of=0&tag=golang"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFeedParamsGetRequestWithDate(t *testing.T) {
	params := NewFeedParams("yukihir0")
	params.SetDate(time.Date(2015, time.August, 1, 12, 34, 56, 0, time.UTC))
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/rss?of=0&date=20150801"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFeedParamsGetRequestWithURL(t *testing.T) {
	params := NewFeedParams("yukihir0")
	params.SetURL("http://d.hatena.ne.jp")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/rss?of=0&url=http://d.hatena.ne.jp"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
