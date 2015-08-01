package hbapi

import (
	"testing"
)

func TestFavoriteFeedParamsGetRequest(t *testing.T) {
	params := NewFavoriteFeedParams("yukihir0")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/favorite.rss?of=0"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
