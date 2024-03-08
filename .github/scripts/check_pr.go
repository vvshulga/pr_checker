package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

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

var SkipLabels = [...]string{"hotfix"}

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

// func normalizeDescription(description string) string {
// 	description = strings.Replace(description, "\r\n", "\n", -1)
// 	description = markdownCommentRegex.ReplaceAllString(description, "")
// 	description = strings.TrimSpace(description)

// 	return description
// }

func main() {
	fmt.Println("Test message - go script was run successfully.")

	cfg := initConfig()

	client := newGithubClient(cfg.token)

	pr, _, _ := client.PullRequests.Get(context.Background(), cfg.repoOwner, cfg.repoName, cfg.prNumber)

	skipCheck := false
	for _, label := range pr.Labels {
		for _, exemptLabel := range SkipLabels {
			if label.GetName() == strings.Trim(exemptLabel, " ") {
				skipCheck = true
				break
			}
		}
	}

	if skipCheck {
		githubactions.Infof("Skipping check because of exempt label")
		os.Exit(0)
	}

	// description := normalizeDescription(pr.GetBody())
	description := pr.GetBody()
	fmt.Println(description)
	githubactions.Infof("TEST message - INFO method")
}
