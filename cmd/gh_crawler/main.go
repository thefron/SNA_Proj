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

	url, err := octokit.UserURL.Expand(octokit.M{"user": "thefron"})

	if err != nil {
		fmt.Println("Error has occurred: ", err)
	}

	user, result := client.Users(url).One()
	if result.HasError() {
		// Handle error
	}

	fmt.Println(user.ReposURL)
}
