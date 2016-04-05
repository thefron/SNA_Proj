package main

import (
	"flag"
	"fmt"

	"github.com/octokit/go-octokit/octokit"
)

func main() {
	token := flag.String("token", "", "github access token")
	flag.Parse()
	var client *octokit.Client
	if *token != "" {
		fmt.Println("Crawling with token:", *token)
		client = octokit.NewClient(octokit.TokenAuth{AccessToken: *token})
	} else {
		fmt.Println("Crawling without token")
		client = octokit.NewClient(nil)
	}

	userchan := GetUser(*client, "thefron")
	user := <-userchan
	fmt.Println(user.ReposURL)
}

func GetUser(client octokit.Client, username string) <-chan octokit.User {
	out := make(chan octokit.User)
	go func() {
		url, err := octokit.UserURL.Expand(octokit.M{"user": username})
		if err != nil {
			fmt.Println("Error has occurred with {}: ", username, err)
		}

		user, result := client.Users(url).One()
		if result.HasError() {
			// TODO: Handle error
		}
		out <- *user
		close(out)
	}()
	return out
}
