package hbapi

// HotEntryFeedParams represents request parameters for get hatena bookmark hot entry feed.
type HotEntryFeedParams struct {
}

// NewHotEntryFeedParams initialize HotEntryFeedParams.
func NewHotEntryFeedParams() HotEntryFeedParams {
	return HotEntryFeedParams{}
}

// GetRequest return request url for get hatena bookmark hot entry feed.
func (params HotEntryFeedParams) GetRequest() string {
	return "https://feeds.feedburner.com/hatena/b/hotentry"
}
