package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/quii/code-review-report/github"
	"github.com/quii/code-review-report/report"
	"log"
	"os"
	"time"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var repos arrayFlags
	var owner string
	flag.Var(&repos, "repo", "name of github repo")
	flag.StringVar(&owner, "owner", "", "owner of repos")
	flag.Parse()

	githubToken := os.Getenv("GITHUB_TOKEN")

	client := github.NewClient(githubToken, os.Stderr)
	service := github.NewService(client)

	var reports []report.IntegrationReport

	log.Printf("Fetching report(s) for %v, owned by %q", repos, owner)

	for _, repo := range repos {
		commits, err := service.GetCommits(context.Background(), Monday(), owner, repo)

		if err != nil {
			log.Fatal(err)
		}
		reports = append(reports, report.NewIntegrationReport(commits, repo))
	}

	json.NewEncoder(os.Stdout).Encode(reports)

}

func Monday() time.Time {
	date := time.Now()

	for date.Weekday() != time.Monday {
		date = date.Add(-1 * (time.Hour * 24))
	}

	year, month, day := date.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, date.Location())
}
