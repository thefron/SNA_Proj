package main

type Organization struct {
	Login            string `json:"login"`
	Id               int    `json:"id"`
	Url              string `json:"url"`
	ReposUrl         string `json:"repos_url"`
	EventsUrl        string `json:"events_url"`
	HooksUrl         string `json:"hooks_url"`
	IssuesUrl        string `json:"issues_url"`
	MembersUrl       string `json:"members_url"`
	PublicMembersUrl string `json:"public_members_url"`
	AvatarUrl        string `json:"avatar_url"`
	Description      string `json:"description"`
}

type OrganizationDetail struct {
	Organization
	Name        *string `json:"name"`
	Company     *string `json:"company"`
	Blog        *string `json:"blog"`
	Location    *string `json:"location"`
	Email       *string `json:"email"`
	PublicRepos int     `json:"public_repos"`
	PublicGists int     `json:"public_gists"`
	Followers   int     `json:"followers"`
	Following   int     `json:"following"`
	HtmlUrl     string  `json:"html_url"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	Type        string  `json:"type"`
}
