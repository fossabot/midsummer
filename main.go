package main

import (
	"fmt"

	"github.com/micnncim/mediumorphose/markdown"
)

func main() {
	fmt.Printf("%+v\n", markdown.CreateSnippets("example.md"))
}
