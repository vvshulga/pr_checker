package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v59/github"
	"github.com/sethvargo/go-githubactions"
	"golang.org/x/oauth2"
)

type config struct {
	token     string
	repoOwner string
	repoName  string
	prNumber  int
}

func initConfig() *config {
	cfg := config{
		token:     githubactions.GetInput("repo-token"),
		repoOwner: githubactions.GetInput("repo-owner"),
		repoName:  githubactions.GetInput("repo-name"),
		prNumber:  strconv.Atoi(githubactions.GetInput("pr-number")),
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

	pr, _, _ := client.PullRequests.Get(context.Background(), cfg.repoOwner, cfg.repoName, cfg.prNumber)

	fmt.Println(pr)
}
