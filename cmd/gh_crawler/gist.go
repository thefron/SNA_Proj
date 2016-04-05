package main

type Gist struct {
	Url         string          `json:"url"`
	ForksUrl    string          `json:"forks_url"`
	CommitsUrl  string          `json:"commits_url"`
	Id          string          `json:"id"`
	GitPullUrl  string          `json:"git_pull_url"`
	GitPushUrl  string          `json:"git_push_url"`
	HtmlUrl     string          `json:"html_url"`
	Files       map[string]File `json:"files,omitempty"`
	Public      bool            `json:"public"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
	Description string          `json:"description"`
	Comments    int             `json:"comments"`
	User        *string         `json:"user"`
	CommentsUrl string          `json:"comments_url"`
	Owner       User            `json:"owner"`
}

type GistDetail struct {
	Gist
	Truncated bool `json:"truncated"`
}
