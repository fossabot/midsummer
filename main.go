package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/micnncim/mediumorphose/config"
	"github.com/micnncim/mediumorphose/gist"
	"github.com/micnncim/mediumorphose/markdown"
	"github.com/micnncim/mediumorphose/medium"
)

var cnf config.Config

func init() {
	if err := cnf.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	g, err := gist.New(cnf.GistConfig.Token)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) != 2 {
		log.Fatal(errors.New("invalid args"))
	}

	md, err := markdown.New(os.Args[1])
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

	mid := medium.New(cnf.MediumConfig.Token)
	if err := mid.Publish(md); err != nil {
		log.Fatal(err)
	}
}
