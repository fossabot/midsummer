package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"github.com/micnncim/mediumorphose/gist"
	"github.com/micnncim/mediumorphose/markdown"
)

func main() {
	g, err := gist.NewClient("09557af82a8d3529b93be7cc7b356cf32c1d9399")
	if err != nil {
		log.Fatal(err)
	}

	snippets := markdown.CreateSnippets("example.md")
	for _, s := range snippets {
		files := map[github.GistFilename]github.GistFile{
			github.GistFilename(s.Filename): github.GistFile{
				Content: &s.Content,
			},
		}
		item, err := g.Create(context.Background(), files, "", true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(*item.HTMLURL)
	}
}
