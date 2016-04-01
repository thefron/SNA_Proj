package main

import (
	"fmt"

	"github.com/octokit/go-octokit/octokit"
)

func main() {
	client := octokit.NewClient(nil)

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
