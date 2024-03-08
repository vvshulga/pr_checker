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
	prNumber, _ := strconv.Atoi(githubactions.GetInput("pr-number"))
	cfg := config{
		token:     githubactions.GetInput("repo-token"),
		repoOwner: githubactions.GetInput("repo-owner"),
		repoName:  githubactions.GetInput("repo-name"),
		prNumber:  prNumber,
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

	pr, a, b := client.PullRequests.Get(context.Background(), cfg.repoOwner, cfg.repoName, cfg.prNumber)

	fmt.Println(pr)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(cfg.repoOwner)
	fmt.Println(cfg.repoName)
	fmt.Println(cfg.prNumber)
}
