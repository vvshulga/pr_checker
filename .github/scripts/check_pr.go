package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v57/github"
	"github.com/sethvargo/go-githubactions"
	"golang.org/x/oauth2"
)

type config struct {
	token string
}

func initConfig() *config {
	cfg := config{
		token: githubactions.GetInput("repo-token"),
	}
	return &cfg
}
func newGithubClient(token string) *github.Client {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func main() {
	fmt.Println("Test message - go script was run successfully.")

	cfg := initConfig()

	client := newGithubClient(cfg.token)

	fmt.Println(client)
}
