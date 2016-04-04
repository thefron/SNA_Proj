package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSgkim126(t *testing.T) {
	// From https://api.github.com/users/sgkim126
	input := `{
  "login": "sgkim126",
  "id": 1138402,
  "avatar_url": "https://avatars.githubusercontent.com/u/1138402?v=3",
  "gravatar_id": "",
  "url": "https://api.github.com/users/sgkim126",
  "html_url": "https://github.com/sgkim126",
  "followers_url": "https://api.github.com/users/sgkim126/followers",
  "following_url": "https://api.github.com/users/sgkim126/following{/other_user}",
  "gists_url": "https://api.github.com/users/sgkim126/gists{/gist_id}",
  "starred_url": "https://api.github.com/users/sgkim126/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/sgkim126/subscriptions",
  "organizations_url": "https://api.github.com/users/sgkim126/orgs",
  "repos_url": "https://api.github.com/users/sgkim126/repos",
  "events_url": "https://api.github.com/users/sgkim126/events{/privacy}",
  "received_events_url": "https://api.github.com/users/sgkim126/received_events",
  "type": "User",
  "site_admin": false,
  "name": "Seulgi Kim",
  "company": null,
  "blog": "http://blog.seulgi.kim",
  "location": "Seoul, Korea",
  "email": null,
  "hireable": null,
  "bio": null,
  "public_repos": 37,
  "public_gists": 111,
  "followers": 23,
  "following": 27,
  "created_at": "2011-10-19T14:35:26Z",
  "updated_at": "2016-04-04T07:38:30Z"
}`
	output := UserDetail{}
	json.Unmarshal([]byte(input), &output)
	assert.Equal(t, "sgkim126", output.Login)
	assert.Equal(t, 1138402, output.Id)
	assert.Equal(t, "https://avatars.githubusercontent.com/u/1138402?v=3", output.AvatarUrl)
	assert.Equal(t, "", output.GravatarId)
	assert.Equal(t, "https://api.github.com/users/sgkim126", output.Url)
	assert.Equal(t, "https://github.com/sgkim126", output.HtmlUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/followers", output.FollowersUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/following{/other_user}", output.FollowingUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/gists{/gist_id}", output.GistsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/starred{/owner}{/repo}", output.StarredUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/subscriptions", output.SubscriptionsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/orgs", output.OrganizationsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/repos", output.ReposUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/events{/privacy}", output.EventsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/received_events", output.ReceivedEventsUrl)
	assert.Equal(t, "User", output.Type)
	assert.False(t, output.SiteAdmin)
	assert.Equal(t, "Seulgi Kim", *output.Name)
	assert.Nil(t, output.Company)
	assert.Equal(t, "http://blog.seulgi.kim", *output.Blog)
	assert.Equal(t, "Seoul, Korea", *output.Location)
	assert.Nil(t, output.Email)
	assert.Nil(t, output.Hireable)
	assert.Nil(t, output.Bio)
	assert.Equal(t, 37, output.PublicRepos)
	assert.Equal(t, 111, output.PublicGists)
	assert.Equal(t, 23, output.Followers)
	assert.Equal(t, 27, output.Following)
	assert.Equal(t, "2011-10-19T14:35:26Z", output.CreatedAt)
	assert.Equal(t, "2016-04-04T07:38:30Z", output.UpdatedAt)
}
