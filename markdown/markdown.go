package markdown

import (
	"bufio"
	"errors"
	"io/ioutil"
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

func (m *Markdown) Parse() error {
	f, err := os.Open(m.Filename)
	if err != nil {
		return err
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

	if len(m.Snippets) == 0 {
		return errors.New("code blocks not exist")
	}
	return nil
}

func (m *Markdown) Replace(urls ...string) error {
	if len(m.Snippets) == 0 {
		return errors.New("code blocks not exist")
	}
	if len(urls) != len(m.Snippets) {
		return errors.New("the number of URLs not match that of code blocks")
	}

	data, err := ioutil.ReadFile(m.Filename)
	if err != nil {
		return err
	}
	content := string(data)

	for i, s := range m.Snippets {
		block := "```" + s.Filename + "\n" + s.Content + "\n```"
		content = strings.Replace(content, block, urls[i], 1)
	}

	if err := ioutil.WriteFile(m.Filename, []byte(content), 0666); err != nil {
		return err
	}
	return nil
}
