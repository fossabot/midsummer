package markdown

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Snippet struct {
	Filename string
	Content  string
}

func CreateSnippets(filename string) []Snippet {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var snippets []Snippet
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "```") {
			continue
		}
		filename := strings.Trim(line, "`")
		var code []string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "```") {
				break
			}
			code = append(code, line)
		}

		for i := range code {
			if i == len(code)-1 {
				break
			}
			code[i] += "\n"
		}

		var content string
		for _, c := range code {
			content += c
		}
		snippets = append(snippets, Snippet{
			Content:  content,
			Filename: filename,
		})
	}

	return snippets
}
