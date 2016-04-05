package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSNA_Proj(t *testing.T) {
	// From https://api.github.com/repos/thefron/SNA_Proj
	input := `{
  "id": 55232227,
  "name": "SNA_Proj",
  "full_name": "thefron/SNA_Proj",
  "owner": {
    "login": "thefron",
    "id": 449334,
    "avatar_url": "https://avatars.githubusercontent.com/u/449334?v=3",
    "gravatar_id": "",
    "url": "https://api.github.com/users/thefron",
    "html_url": "https://github.com/thefron",
    "followers_url": "https://api.github.com/users/thefron/followers",
    "following_url": "https://api.github.com/users/thefron/following{/other_user}",
    "gists_url": "https://api.github.com/users/thefron/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/thefron/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/thefron/subscriptions",
    "organizations_url": "https://api.github.com/users/thefron/orgs",
    "repos_url": "https://api.github.com/users/thefron/repos",
    "events_url": "https://api.github.com/users/thefron/events{/privacy}",
    "received_events_url": "https://api.github.com/users/thefron/received_events",
    "type": "User",
    "site_admin": false
  },
  "private": false,
  "html_url": "https://github.com/thefron/SNA_Proj",
  "description": null,
  "fork": false,
  "url": "https://api.github.com/repos/thefron/SNA_Proj",
  "forks_url": "https://api.github.com/repos/thefron/SNA_Proj/forks",
  "keys_url": "https://api.github.com/repos/thefron/SNA_Proj/keys{/key_id}",
  "collaborators_url": "https://api.github.com/repos/thefron/SNA_Proj/collaborators{/collaborator}",
  "teams_url": "https://api.github.com/repos/thefron/SNA_Proj/teams",
  "hooks_url": "https://api.github.com/repos/thefron/SNA_Proj/hooks",
  "issue_events_url": "https://api.github.com/repos/thefron/SNA_Proj/issues/events{/number}",
  "events_url": "https://api.github.com/repos/thefron/SNA_Proj/events",
  "assignees_url": "https://api.github.com/repos/thefron/SNA_Proj/assignees{/user}",
  "branches_url": "https://api.github.com/repos/thefron/SNA_Proj/branches{/branch}",
  "tags_url": "https://api.github.com/repos/thefron/SNA_Proj/tags",
  "blobs_url": "https://api.github.com/repos/thefron/SNA_Proj/git/blobs{/sha}",
  "git_tags_url": "https://api.github.com/repos/thefron/SNA_Proj/git/tags{/sha}",
  "git_refs_url": "https://api.github.com/repos/thefron/SNA_Proj/git/refs{/sha}",
  "trees_url": "https://api.github.com/repos/thefron/SNA_Proj/git/trees{/sha}",
  "statuses_url": "https://api.github.com/repos/thefron/SNA_Proj/statuses/{sha}",
  "languages_url": "https://api.github.com/repos/thefron/SNA_Proj/languages",
  "stargazers_url": "https://api.github.com/repos/thefron/SNA_Proj/stargazers",
  "contributors_url": "https://api.github.com/repos/thefron/SNA_Proj/contributors",
  "subscribers_url": "https://api.github.com/repos/thefron/SNA_Proj/subscribers",
  "subscription_url": "https://api.github.com/repos/thefron/SNA_Proj/subscription",
  "commits_url": "https://api.github.com/repos/thefron/SNA_Proj/commits{/sha}",
  "git_commits_url": "https://api.github.com/repos/thefron/SNA_Proj/git/commits{/sha}",
  "comments_url": "https://api.github.com/repos/thefron/SNA_Proj/comments{/number}",
  "issue_comment_url": "https://api.github.com/repos/thefron/SNA_Proj/issues/comments{/number}",
  "contents_url": "https://api.github.com/repos/thefron/SNA_Proj/contents/{+path}",
  "compare_url": "https://api.github.com/repos/thefron/SNA_Proj/compare/{base}...{head}",
  "merges_url": "https://api.github.com/repos/thefron/SNA_Proj/merges",
  "archive_url": "https://api.github.com/repos/thefron/SNA_Proj/{archive_format}{/ref}",
  "downloads_url": "https://api.github.com/repos/thefron/SNA_Proj/downloads",
  "issues_url": "https://api.github.com/repos/thefron/SNA_Proj/issues{/number}",
  "pulls_url": "https://api.github.com/repos/thefron/SNA_Proj/pulls{/number}",
  "milestones_url": "https://api.github.com/repos/thefron/SNA_Proj/milestones{/number}",
  "notifications_url": "https://api.github.com/repos/thefron/SNA_Proj/notifications{?since,all,participating}",
  "labels_url": "https://api.github.com/repos/thefron/SNA_Proj/labels{/name}",
  "releases_url": "https://api.github.com/repos/thefron/SNA_Proj/releases{/id}",
  "deployments_url": "https://api.github.com/repos/thefron/SNA_Proj/deployments",
  "created_at": "2016-04-01T12:54:18Z",
  "updated_at": "2016-04-01T14:53:53Z",
  "pushed_at": "2016-04-01T12:54:31Z",
  "git_url": "git://github.com/thefron/SNA_Proj.git",
  "ssh_url": "git@github.com:thefron/SNA_Proj.git",
  "clone_url": "https://github.com/thefron/SNA_Proj.git",
  "svn_url": "https://github.com/thefron/SNA_Proj",
  "homepage": null,
  "size": 50,
  "stargazers_count": 1,
  "watchers_count": 1,
  "language": "Go",
  "has_issues": true,
  "has_downloads": true,
  "has_wiki": true,
  "has_pages": false,
  "forks_count": 1,
  "mirror_url": null,
  "open_issues_count": 0,
  "forks": 1,
  "open_issues": 0,
  "watchers": 1,
  "default_branch": "master",
  "network_count": 1,
  "subscribers_count": 2
}`
	output := Repository{}
	json.Unmarshal([]byte(input), &output)
	assert.Equal(t, 55232227, output.Id)
	assert.Equal(t, "SNA_Proj", output.Name)
	assert.Equal(t, "thefron/SNA_Proj", output.FullName)
	assert.Equal(t, "thefron", output.Owner.Login)
	assert.Equal(t, 449334, output.Owner.Id)
	assert.Equal(t, "https://avatars.githubusercontent.com/u/449334?v=3", output.Owner.AvatarUrl)
	assert.Equal(t, "", output.Owner.GravatarId)
	assert.Equal(t, "https://api.github.com/users/thefron", output.Owner.Url)
	assert.Equal(t, "https://github.com/thefron", output.Owner.HtmlUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/followers", output.Owner.FollowersUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/following{/other_user}", output.Owner.FollowingUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/gists{/gist_id}", output.Owner.GistsUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/starred{/owner}{/repo}", output.Owner.StarredUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/subscriptions", output.Owner.SubscriptionsUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/orgs", output.Owner.OrganizationsUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/repos", output.Owner.ReposUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/events{/privacy}", output.Owner.EventsUrl)
	assert.Equal(t, "https://api.github.com/users/thefron/received_events", output.Owner.ReceivedEventsUrl)
	assert.Equal(t, "User", output.Owner.Type)
	assert.False(t, output.Owner.SiteAdmin)
	assert.False(t, output.Private)
	assert.Equal(t, "https://github.com/thefron/SNA_Proj", output.HtmlUrl)
	assert.Nil(t, output.Description)
	assert.False(t, output.Fork)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj", output.Url)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/forks", output.ForksUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/keys{/key_id}", output.KeysUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/collaborators{/collaborator}", output.CollaboratorsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/teams", output.TeamsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/hooks", output.HooksUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/issues/events{/number}", output.IssueEventsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/events", output.EventsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/assignees{/user}", output.AssigneesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/branches{/branch}", output.BranchesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/tags", output.TagsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/git/blobs{/sha}", output.BlobsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/git/tags{/sha}", output.GitTagsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/git/refs{/sha}", output.GitRefsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/git/trees{/sha}", output.TreesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/statuses/{sha}", output.StatusesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/languages", output.LanguagesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/stargazers", output.StargazersUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/contributors", output.ContributorsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/subscribers", output.SubscribersUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/subscription", output.SubscriptionUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/commits{/sha}", output.CommitsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/git/commits{/sha}", output.GitCommitsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/comments{/number}", output.CommentsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/issues/comments{/number}", output.IssueCommentUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/contents/{+path}", output.ContentsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/compare/{base}...{head}", output.CompareUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/merges", output.MergesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/{archive_format}{/ref}", output.ArchiveUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/downloads", output.DownloadsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/issues{/number}", output.IssuesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/pulls{/number}", output.PullsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/milestones{/number}", output.MilestonesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/notifications{?since,all,participating}", output.NotificationsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/labels{/name}", output.LabelsUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/releases{/id}", output.ReleasesUrl)
	assert.Equal(t, "https://api.github.com/repos/thefron/SNA_Proj/deployments", output.DeploymentsUrl)
	assert.Equal(t, "2016-04-01T12:54:18Z", output.CreatedAt)
	assert.Equal(t, "2016-04-01T14:53:53Z", output.UpdatedAt)
	assert.Equal(t, "2016-04-01T12:54:31Z", output.PushedAt)
	assert.Equal(t, "git://github.com/thefron/SNA_Proj.git", output.GitUrl)
	assert.Equal(t, "git@github.com:thefron/SNA_Proj.git", output.SshUrl)
	assert.Equal(t, "https://github.com/thefron/SNA_Proj.git", output.CloneUrl)
	assert.Equal(t, "https://github.com/thefron/SNA_Proj", output.SvnUrl)
	assert.Nil(t, output.Homepage)
	assert.Equal(t, 50, output.Size)
	assert.Equal(t, 1, output.StargazersCount)
	assert.Equal(t, 1, output.WatchersCount)
	assert.Equal(t, "Go", output.Language)
	assert.True(t, output.HasIssues)
	assert.True(t, output.HasDownloads)
	assert.True(t, output.HasWiki)
	assert.False(t, output.HasPages)
	assert.Equal(t, 1, output.ForksCount)
	assert.Nil(t, output.MirrorUrl)
	assert.Equal(t, 0, output.OpenIssuesCount)
	assert.Equal(t, 1, output.Forks)
	assert.Equal(t, 0, output.OpenIssues)
	assert.Equal(t, 1, output.Watchers)
	assert.Equal(t, "master", output.DefaultBranch)
	assert.Equal(t, 1, output.NetworkCount)
	assert.Equal(t, 2, output.SubscribersCount)

}
