package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"
)

func main() {
	token := "GITHUB_PERSONAL_ACCESS_TOKEN"
	username := "byigitt"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.ListOptions{
		PerPage: 100,
		Page:    1,
	}

	var allFollowing []*github.User
	for {
		following, resp, err := client.Users.ListFollowing(ctx, username, opt)
		if err != nil {
			log.Fatal(err)
		}
		allFollowing = append(allFollowing, following...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, user := range allFollowing {
		isFollowing, _, err := client.Users.IsFollowing(ctx, user.GetLogin(), username)
		if err != nil {
			log.Fatal(err)
		}

		if !isFollowing {
			_, err := client.Users.Unfollow(ctx, user.GetLogin())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Unfollowed: %s\n", user.GetLogin())	
		}
	}
}