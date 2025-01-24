package models

type TimeEntry struct {
	EntryType string `json:"entry_type"`
	Worker    int    `json:"worker"`
}
