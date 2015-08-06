package hbapi

import "testing"

func TestHotEntryFeedParamsGetRequest(t *testing.T) {
	params := NewHotEntryFeedParams()
	actual := params.GetRequest()
	expected := "https://feeds.feedburner.com/hatena/b/hotentry"

	if actual != expected {
		t.Errorf("expected %s, but got %s\n", expected, actual)
	}
}
