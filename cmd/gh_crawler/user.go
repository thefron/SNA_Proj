package main

type User struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type UserDetail struct {
	User
	Name        *string `json:"name"`
	Company     *string `json:"company"`
	Blog        *string `json:"blog"`
	Location    *string `json:"location"`
	Email       *string `json:"email"`
	Hireable    *bool   `json:"hireable"`
	Bio         *string `json:"bio"`
	PublicRepos int     `json:"public_repos"`
	PublicGists int     `json:"public_gists"`
	Followers   int     `json:"followers"`
	Following   int     `json:"following"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
