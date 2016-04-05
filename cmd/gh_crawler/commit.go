package main

type ChangeStatus struct {
	Total     int `json:"total"`
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
}

type Commit struct {
	User         User         `json:"user"`
	Version      string       `json:"version"`
	CommittedAt  string       `json:"committed_at"`
	ChangeStatus ChangeStatus `json:"change_status"`
	Url          string       `json:"url"`
}
