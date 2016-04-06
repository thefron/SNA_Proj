package main

import (
	"github.com/octokit/go-octokit/octokit"
)

func GetFollowers(client octokit.Client, user octokit.User) <-chan string {
	return getFollowers(client, user.FollowersURL)
}

func GetFollowings(client octokit.Client, user octokit.User) <-chan string {
	return getFollowers(client, user.FollowingURL)
}

func getFollowers(client octokit.Client, uri octokit.Hyperlink) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		followers, result := client.Followers().All(&uri, octokit.M{})
		if result.HasError() {
			// TODO: Handle error
		}
		for _, follower := range followers {
			out <- follower.Login
		}
	}()
	return out
}
