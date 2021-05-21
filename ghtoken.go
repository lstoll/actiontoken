package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
)

// https://stackoverflow.com/questions/9765453/is-gits-semi-secret-empty-tree-object-reliable-and-why-is-there-not-a-symbolic
const emptyTreeSHA = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"

func validateToken(ctx context.Context, owner, repo, token string) (bool, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	commit, resp, err := client.Git.CreateCommit(ctx, owner, repo, &github.Commit{
		Message: strPtr("commit by actiontoken to verify write access"),
		Tree: &github.Tree{
			SHA: strPtr(emptyTreeSHA),
		},
		Parents: []*github.Commit{},
	})
	if err != nil {
		if resp.Response != nil && resp.Response.StatusCode == http.StatusUnauthorized {
			return false, nil
		}
		return false, fmt.Errorf("failed creating commit: %v", err)
	}
	// TODO - validate commit
	// i.e make sure it's on the expected repo
	// check that the signature was by github?
	_ = commit

	return true, nil
}

func strPtr(s string) *string {
	return &s
}
