package hbapi

// entry record in hatena bookmark entry information
type HBEntry struct {
	Title    string `json:"title"`
	Count    int    `json:"count"`
	Url      string `json:"url"`
	EntryUrl string `json:"entry_url"`
	Eid      int    `json:"eid"`
}

// bookmark record in hatena bookmark entry information
type HBBookmark struct {
	User      string   `json:"user"`
	Tags      []string `json:"tags"`
	Timestamp string   `json:"timestamp"`
	Comment   string   `json:"comment"`
}

// harena bookmark entry information
type HBEntryInfo struct {
	HBEntry
	Screenshot string       `json:"screenshot"`
	Bookmarks  []HBBookmark `json:"bookmarks"`
	Related    []HBEntry    `json:"related"`
}
