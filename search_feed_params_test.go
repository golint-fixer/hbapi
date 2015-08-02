package hbapi

import (
	"testing"
)

func TestSearchFeedParamsGetRequest(t *testing.T) {
	params := NewSearchFeedParams("golang")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/search/text?q=golang&sort=recent&users=3&safe=on&mode=rss"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
