package github

import (
	"context"
	"time"
)

type SimpleCommit struct {
	Email        string
	AvatarURL    string
	Message      string
	Status       string
	CreatedAt    time.Time
	FilesChanged []string
}

func (s SimpleCommit) Successful() bool {
	return s.Status == "success"
}

func (s SimpleCommit) Failed() bool {
	return s.Status == "failure"
}

type CommitService interface {
	GetCommits(ctx context.Context, since time.Time, owner string, repos ...string) ([]SimpleCommit, error)
}

type AliasService interface {
	GetAlias(string) string
}
