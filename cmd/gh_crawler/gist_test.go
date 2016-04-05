package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSgkim126Gists(t *testing.T) {
	// From https://api.github.com/users/sgkim126/gists
	input := `{
    "url": "https://api.github.com/gists/e067eea45bc973f41b5c",
    "forks_url": "https://api.github.com/gists/e067eea45bc973f41b5c/forks",
    "commits_url": "https://api.github.com/gists/e067eea45bc973f41b5c/commits",
    "id": "e067eea45bc973f41b5c",
    "git_pull_url": "https://gist.github.com/e067eea45bc973f41b5c.git",
    "git_push_url": "https://gist.github.com/e067eea45bc973f41b5c.git",
    "html_url": "https://gist.github.com/e067eea45bc973f41b5c",
    "files": {
      "SomeTask.java": {
        "filename": "SomeTask.java",
        "type": "text/plain",
        "language": "Java",
        "raw_url": "https://gist.githubusercontent.com/sgkim126/e067eea45bc973f41b5c/raw/7080b5d42f2d7dd9fb5b2cc1a2df49dec6b38017/SomeTask.java",
        "size": 377
      },
      "main.java": {
        "filename": "main.java",
        "type": "text/plain",
        "language": "Java",
        "raw_url": "https://gist.githubusercontent.com/sgkim126/e067eea45bc973f41b5c/raw/39e4b56a011a365784b50eab5b39fbb63a16d29a/main.java",
        "size": 53
      }
    },
    "public": true,
    "created_at": "2016-03-06T15:56:16Z",
    "updated_at": "2016-03-06T16:02:05Z",
    "description": "",
    "comments": 0,
    "user": null,
    "comments_url": "https://api.github.com/gists/e067eea45bc973f41b5c/comments",
    "owner": {
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
      "site_admin": false
    },
    "truncated": false
  }`
	output := GistDetail{}
	json.Unmarshal([]byte(input), &output)

	assert.Equal(t, "https://api.github.com/gists/e067eea45bc973f41b5c", output.Url)
	assert.Equal(t, "https://api.github.com/gists/e067eea45bc973f41b5c/forks", output.ForksUrl)
	assert.Equal(t, "https://api.github.com/gists/e067eea45bc973f41b5c/commits", output.CommitsUrl)
	assert.Equal(t, "e067eea45bc973f41b5c", output.Id)
	assert.Equal(t, "https://gist.github.com/e067eea45bc973f41b5c.git", output.GitPullUrl)
	assert.Equal(t, "https://gist.github.com/e067eea45bc973f41b5c.git", output.GitPushUrl)
	assert.Equal(t, "https://gist.github.com/e067eea45bc973f41b5c", output.HtmlUrl)
	assert.Equal(t, "SomeTask.java", output.Files["SomeTask.java"].Filename)
	assert.Equal(t, "text/plain", output.Files["SomeTask.java"].Type)
	assert.Equal(t, "Java", output.Files["SomeTask.java"].Language)
	assert.Equal(t, "https://gist.githubusercontent.com/sgkim126/e067eea45bc973f41b5c/raw/7080b5d42f2d7dd9fb5b2cc1a2df49dec6b38017/SomeTask.java", output.Files["SomeTask.java"].RawUrl)
	assert.Equal(t, 377, output.Files["SomeTask.java"].Size)
	assert.Equal(t, "main.java", output.Files["main.java"].Filename)
	assert.Equal(t, "text/plain", output.Files["main.java"].Type)
	assert.Equal(t, "Java", output.Files["main.java"].Language)
	assert.Equal(t, "https://gist.githubusercontent.com/sgkim126/e067eea45bc973f41b5c/raw/39e4b56a011a365784b50eab5b39fbb63a16d29a/main.java", output.Files["main.java"].RawUrl)
	assert.Equal(t, 53, output.Files["main.java"].Size)
	assert.True(t, output.Public)
	assert.Equal(t, "2016-03-06T15:56:16Z", output.CreatedAt)
	assert.Equal(t, "2016-03-06T16:02:05Z", output.UpdatedAt)
	assert.Equal(t, "", output.Description)
	assert.Equal(t, 0, output.Comments)
	assert.Nil(t, output.User)
	assert.Equal(t, "https://api.github.com/gists/e067eea45bc973f41b5c/comments", output.CommentsUrl)
	assert.Equal(t, "sgkim126", output.Owner.Login)
	assert.Equal(t, 1138402, output.Owner.Id)
	assert.Equal(t, "https://avatars.githubusercontent.com/u/1138402?v=3", output.Owner.AvatarUrl)
	assert.Equal(t, "", output.Owner.GravatarId)
	assert.Equal(t, "https://api.github.com/users/sgkim126", output.Owner.Url)
	assert.Equal(t, "https://github.com/sgkim126", output.Owner.HtmlUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/followers", output.Owner.FollowersUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/following{/other_user}", output.Owner.FollowingUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/gists{/gist_id}", output.Owner.GistsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/starred{/owner}{/repo}", output.Owner.StarredUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/subscriptions", output.Owner.SubscriptionsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/orgs", output.Owner.OrganizationsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/repos", output.Owner.ReposUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/events{/privacy}", output.Owner.EventsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/received_events", output.Owner.ReceivedEventsUrl)
	assert.Equal(t, "User", output.Owner.Type)
	assert.False(t, output.Owner.SiteAdmin)
	assert.False(t, output.Truncated)
}
