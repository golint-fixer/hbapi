package hbapi

import (
	"testing"
)

func TestFeedParamsGetRequest(t *testing.T) {
	params := NewFeedParams("yukihir0")
	actual := params.GetRequest()
	expected := "http://b.hatena.ne.jp/yukihir0/rss?of=0"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
