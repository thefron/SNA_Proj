package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithub(t *testing.T) {
	// From https://api.github.com/orgs/github
	input := `{
  "login": "github",
  "id": 9919,
  "url": "https://api.github.com/orgs/github",
  "repos_url": "https://api.github.com/orgs/github/repos",
  "events_url": "https://api.github.com/orgs/github/events",
  "hooks_url": "https://api.github.com/orgs/github/hooks",
  "issues_url": "https://api.github.com/orgs/github/issues",
  "members_url": "https://api.github.com/orgs/github/members{/member}",
  "public_members_url": "https://api.github.com/orgs/github/public_members{/member}",
  "avatar_url": "https://avatars.githubusercontent.com/u/9919?v=3",
  "description": "How people build software.",
  "name": "GitHub",
  "company": null,
  "blog": "https://github.com/about",
  "location": "San Francisco, CA",
  "email": "support@github.com",
  "public_repos": 128,
  "public_gists": 0,
  "followers": 0,
  "following": 0,
  "html_url": "https://github.com/github",
  "created_at": "2008-05-11T04:37:31Z",
  "updated_at": "2016-03-15T22:48:02Z",
  "type": "Organization"
}`
	output := OrganizationDetail{}
	json.Unmarshal([]byte(input), &output)
	assert.Equal(t, "github", output.Login)
	assert.Equal(t, 9919, output.Id)
	assert.Equal(t, "https://api.github.com/orgs/github", output.Url)
	assert.Equal(t, "https://api.github.com/orgs/github/repos", output.ReposUrl)
	assert.Equal(t, "https://api.github.com/orgs/github/events", output.EventsUrl)
	assert.Equal(t, "https://api.github.com/orgs/github/hooks", output.HooksUrl)
	assert.Equal(t, "https://api.github.com/orgs/github/issues", output.IssuesUrl)
	assert.Equal(t, "https://api.github.com/orgs/github/members{/member}", output.MembersUrl)
	assert.Equal(t, "https://api.github.com/orgs/github/public_members{/member}", output.PublicMembersUrl)
	assert.Equal(t, "https://avatars.githubusercontent.com/u/9919?v=3", output.AvatarUrl)
	assert.Equal(t, "How people build software.", output.Description)
	assert.Equal(t, "GitHub", *output.Name)
	assert.Nil(t, output.Company)
	assert.Equal(t, "https://github.com/about", *output.Blog)
	assert.Equal(t, "San Francisco, CA", *output.Location)
	assert.Equal(t, "support@github.com", *output.Email)
	assert.Equal(t, 128, output.PublicRepos)
	assert.Equal(t, 0, output.PublicGists)
	assert.Equal(t, 0, output.Followers)
	assert.Equal(t, 0, output.Following)
	assert.Equal(t, "https://github.com/github", output.HtmlUrl)
	assert.Equal(t, "2008-05-11T04:37:31Z", output.CreatedAt)
	assert.Equal(t, "2016-03-15T22:48:02Z", output.UpdatedAt)
	assert.Equal(t, "Organization", output.Type)
}
