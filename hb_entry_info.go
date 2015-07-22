package hbapi

// HBEntry represents entry record.
type HBEntry struct {
	Title    string `json:"title"`
	Count    int    `json:"count"`
	URL      string `json:"url"`
	EntryURL string `json:"entry_url"`
	Eid      int    `json:"eid"`
}

// HBBookmark represents bookamrk record.
type HBBookmark struct {
	User      string   `json:"user"`
	Tags      []string `json:"tags"`
	Timestamp string   `json:"timestamp"`
	Comment   string   `json:"comment"`
}

// HBEntryInfo represents entry info record.
type HBEntryInfo struct {
	HBEntry
	Screenshot string       `json:"screenshot"`
	Bookmarks  []HBBookmark `json:"bookmarks"`
	Related    []HBEntry    `json:"related"`
}
