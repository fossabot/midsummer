package gist

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Gist struct {
	Client *github.Client
}

func NewClient(token string) (*Gist, error) {
	if token == "" {
		return &Gist{}, errors.New("token is missing")
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	return &Gist{Client: github.NewClient(tc)}, nil
}

func (g *Gist) Create(
	ctx context.Context,
	files map[github.GistFilename]github.GistFile,
	desc string,
	public bool) (*github.Gist, error) {

	item, resp, err := g.Client.Gists.Create(ctx, &github.Gist{
		Files:       files,
		Description: &desc,
		Public:      &public,
	})
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("gist item is nil")
	}
	if resp == nil {
		return nil, errors.New("resp is nil")
	}

	return item, nil
}
