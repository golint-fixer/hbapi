package hbapi

import (
	"testing"
	"time"
)

func TestFavoriteFeedParamsGetRequest(t *testing.T) {
	params := NewFavoriteFeedParams("yukihir0")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/favorite.rss?of=0"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFavoriteFeedParamsGetRequestWithUser(t *testing.T) {
	params := NewFavoriteFeedParams("yukihir0")
	params.SetUser("sample")

	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/sample/favorite.rss?of=0"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFavoriteFeedParamsGetRequestWithPage(t *testing.T) {
	params := NewFavoriteFeedParams("yukihir0")
	params.SetPage(1)

	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/favorite.rss?of=25"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFavoriteFeedParamsGetRequestWithUntil(t *testing.T) {
	params := NewFavoriteFeedParams("yukihir0")
	params.SetUntil(time.Date(2015, time.August, 1, 12, 34, 56, 0, time.UTC))

	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/favorite.rss?until=1438432496"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFavoriteFeedParamsGetRequestWithEnableWithMe(t *testing.T) {
	params := NewFavoriteFeedParams("yukihir0")
	params.EnableWithMe()

	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/favorite.rss?of=0&with_me=1"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}

func TestFavoriteFeedParamsGetRequestWithDisableWithMe(t *testing.T) {
	params := NewFavoriteFeedParams("yukihir0")
	params.DisableWithMe()

	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/favorite.rss?of=0"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
