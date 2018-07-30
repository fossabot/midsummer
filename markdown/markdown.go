package markdown

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type (
	Markdown struct {
		Filename string
		Snippets []*Snippet
	}
	Snippet struct {
		Filename string
		Content  string
	}
)

func New(filename string) *Markdown {
	return &Markdown{Filename: filename}
}

func (m *Markdown) Parse() {
	f, err := os.Open(m.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

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
		m.Snippets = append(m.Snippets, &Snippet{
			Content:  content,
			Filename: filename,
		})
	}
}
