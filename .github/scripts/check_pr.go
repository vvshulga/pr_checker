package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
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

var (
	markdownCommentRegex = regexp.MustCompile(`\<\!\-\-\-.*\-\-\>`)
	blockHeaderRegex     = `(?m)^\s*##([^#].*?)$`
	skipLabels           = [...]string{"hotfix"}
)

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

func normalizeDescription(description string) string {
	description = strings.Replace(description, "\r\n", "\n", -1)
	description = markdownCommentRegex.ReplaceAllString(description, "")
	description = strings.TrimSpace(description)

	return description
}

func splitByBlocks(txt string) []string {
	re := regexp.MustCompile(blockHeaderRegex)

	split := re.Split(txt, -1)
	blocks := []string{}

	for i := range split {
		blocks = append(blocks, split[i])
	}

	fmt.Println(blocks)
	return blocks
}

func main() {
	fmt.Println("Test message - go script was run successfully.")

	cfg := initConfig()

	client := newGithubClient(cfg.token)

	pr, _, _ := client.PullRequests.Get(context.Background(), cfg.repoOwner, cfg.repoName, cfg.prNumber)

	skipCheck := false
	for _, label := range pr.Labels {
		for _, exemptLabel := range skipLabels {
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

	description := normalizeDescription(pr.GetBody())
	fmt.Println(description)

	splitByBlocks(description)
	githubactions.Infof("TEST message - INFO method")
}
