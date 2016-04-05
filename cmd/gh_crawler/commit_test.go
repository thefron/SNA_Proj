package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommit_e067eea45bc973f41b5c(t *testing.T) {
	// From https://api.github.com/gists/e067eea45bc973f41b5c/commits
	input := `[
  {
    "user": {
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
    "version": "3571908a87a85ef45272e7181cdfe4fbb738cebe",
    "committed_at": "2016-03-06T16:02:05Z",
    "change_status": {
      "total": 2,
      "additions": 1,
      "deletions": 1
    },
    "url": "https://api.github.com/gists/e067eea45bc973f41b5c/3571908a87a85ef45272e7181cdfe4fbb738cebe"
  },
  {
    "user": {
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
    "version": "21b68881d308a146bbb6ffe18ccfa5d7df7703f6",
    "committed_at": "2016-03-06T15:56:16Z",
    "change_status": {
      "total": 19,
      "additions": 19,
      "deletions": 0
    },
    "url": "https://api.github.com/gists/e067eea45bc973f41b5c/21b68881d308a146bbb6ffe18ccfa5d7df7703f6"
  }
]`
	output := make([]Commit, 0)
	json.Unmarshal([]byte(input), &output)
	assert.Equal(t, "sgkim126", output[0].User.Login)
	assert.Equal(t, 1138402, output[0].User.Id)
	assert.Equal(t, "https://avatars.githubusercontent.com/u/1138402?v=3", output[0].User.AvatarUrl)
	assert.Equal(t, "", output[0].User.GravatarId)
	assert.Equal(t, "https://api.github.com/users/sgkim126", output[0].User.Url)
	assert.Equal(t, "https://github.com/sgkim126", output[0].User.HtmlUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/followers", output[0].User.FollowersUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/following{/other_user}", output[0].User.FollowingUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/gists{/gist_id}", output[0].User.GistsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/starred{/owner}{/repo}", output[0].User.StarredUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/subscriptions", output[0].User.SubscriptionsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/orgs", output[0].User.OrganizationsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/repos", output[0].User.ReposUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/events{/privacy}", output[0].User.EventsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/received_events", output[0].User.ReceivedEventsUrl)
	assert.Equal(t, "User", output[0].User.Type)
	assert.False(t, output[0].User.SiteAdmin)

	assert.Equal(t, "3571908a87a85ef45272e7181cdfe4fbb738cebe", output[0].Version)
	assert.Equal(t, "2016-03-06T16:02:05Z", output[0].CommittedAt)
	assert.Equal(t, 2, output[0].ChangeStatus.Total)
	assert.Equal(t, 1, output[0].ChangeStatus.Additions)
	assert.Equal(t, 1, output[0].ChangeStatus.Deletions)
	assert.Equal(t, "https://api.github.com/gists/e067eea45bc973f41b5c/3571908a87a85ef45272e7181cdfe4fbb738cebe", output[0].Url)
	assert.Equal(t, "sgkim126", output[1].User.Login)
	assert.Equal(t, 1138402, output[1].User.Id)
	assert.Equal(t, "https://avatars.githubusercontent.com/u/1138402?v=3", output[1].User.AvatarUrl)
	assert.Equal(t, "", output[1].User.GravatarId)
	assert.Equal(t, "https://api.github.com/users/sgkim126", output[1].User.Url)
	assert.Equal(t, "https://github.com/sgkim126", output[1].User.HtmlUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/followers", output[1].User.FollowersUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/following{/other_user}", output[1].User.FollowingUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/gists{/gist_id}", output[1].User.GistsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/starred{/owner}{/repo}", output[1].User.StarredUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/subscriptions", output[1].User.SubscriptionsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/orgs", output[1].User.OrganizationsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/repos", output[1].User.ReposUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/events{/privacy}", output[1].User.EventsUrl)
	assert.Equal(t, "https://api.github.com/users/sgkim126/received_events", output[1].User.ReceivedEventsUrl)
	assert.Equal(t, "User", output[1].User.Type)
	assert.False(t, output[1].User.SiteAdmin)

	assert.Equal(t, "21b68881d308a146bbb6ffe18ccfa5d7df7703f6", output[1].Version)
	assert.Equal(t, "2016-03-06T15:56:16Z", output[1].CommittedAt)
	assert.Equal(t, 19, output[1].ChangeStatus.Total)
	assert.Equal(t, 19, output[1].ChangeStatus.Additions)
	assert.Equal(t, 0, output[1].ChangeStatus.Deletions)
	assert.Equal(t, "https://api.github.com/gists/e067eea45bc973f41b5c/21b68881d308a146bbb6ffe18ccfa5d7df7703f6", output[1].Url)
}
