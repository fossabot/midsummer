package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/micnncim/mediumorphose/gist"
	"github.com/micnncim/mediumorphose/markdown"
	"github.com/micnncim/mediumorphose/medium"
)

func main() {
	g, err := gist.New(os.Getenv("GITHUB_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	md, err := markdown.New("example.md")
	if err != nil {
		log.Fatal(err)
	}
	if err := md.ParseTitle(); err != nil {
		log.Fatal(err)
	}
	if err := md.ParseSnippets(); err != nil {
		log.Fatal(err)
	}

	var urls []string
	for _, s := range md.Snippets {
		files := map[github.GistFilename]github.GistFile{
			github.GistFilename(s.Filename): github.GistFile{
				Content: &s.Content,
			},
		}
		item, err := g.Create(context.Background(), files, "", true)
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, *item.HTMLURL)
		fmt.Println(*item.HTMLURL)
	}

	if err := md.Replace(urls...); err != nil {
		log.Fatal(err)
	}
	if err := md.Write(); err != nil {
		log.Fatal(err)
	}

	mid := medium.New(os.Getenv("MEDIUM_ACCESS_TOKEN"))
	if err := mid.Publish(md); err != nil {
		log.Fatal(err)
	}
}
