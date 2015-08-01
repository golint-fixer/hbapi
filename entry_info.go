package hbapi

// EntryInfo represents entry info record.
type EntryInfo struct {
	Entry
	Screenshot string     `json:"screenshot"`
	Bookmarks  []Bookmark `json:"bookmarks"`
	Related    []Entry    `json:"related"`
}

// Entry represents entry record.
type Entry struct {
	Title    string `json:"title"`
	Count    int    `json:"count"`
	URL      string `json:"url"`
	EntryURL string `json:"entry_url"`
	Eid      int    `json:"eid"`
}

// Bookmark represents bookamrk record.
type Bookmark struct {
	User      string   `json:"user"`
	Tags      []string `json:"tags"`
	Timestamp string   `json:"timestamp"`
	Comment   string   `json:"comment"`
}
