package report

import (
	"github.com/quii/code-review-report/github"
	"sort"
)

func NewIntegrationReport(commits []github.SimpleCommit, repo string) IntegrationReport {
	auctionsReport := IntegrationReport{Repo: repo}

	changedFiles := make(map[string]int)
	filesThatFailed := make(map[string]int)

	for _, commit := range commits {
		if commit.Successful() {
			auctionsReport.Integrations.Successful++
		}
		if commit.Failed() {
			auctionsReport.Integrations.Failed++
			auctionsReport.FailedCommits = append(auctionsReport.FailedCommits, commit.Message)
		}

		for _, file := range commit.FilesChanged {
			changedFiles[file]++
			if commit.Failed() {
				filesThatFailed[file]++
			}
		}
	}

	auctionsReport.MostChanged = NewFileFrequencies(changedFiles)[:10]
	auctionsReport.FlakyFiles = NewFileFrequencies(filesThatFailed)
	return auctionsReport
}

type IntegrationReport struct {
	Repo         string `json:"repo"`
	Integrations struct {
		Successful int `json:"successful"`
		Failed     int `json:"failed"`
	} `json:"integrations"`
	MostChanged   []FileFrequency `json:"most_changed"`
	FlakyFiles    []FileFrequency `json:"flaky_files"`
	FailedCommits []string        `json:"failed_commits"`
}

type FileFrequency struct {
	Name  string
	Count int
}

func NewFileFrequencies(stats map[string]int) []FileFrequency {
	var f []FileFrequency
	for file, count := range stats {
		f = append(f, FileFrequency{
			Name:  file,
			Count: count,
		})
	}
	sort.Slice(f, func(i, j int) bool {
		return f[i].Count > f[j].Count
	})
	return f
}
