package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/micnncim/mediumorphose/gist"
	"github.com/micnncim/mediumorphose/markdown"
)

func main() {
	token := os.Getenv("GITHUB_ACCESS_TOKEN")
	g, err := gist.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	m := markdown.New("example.md")
	m.Parse()
	for _, s := range m.Snippets {
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
