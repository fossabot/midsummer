package medium

import (
	medium "github.com/Medium/medium-sdk-go"
	"github.com/micnncim/mediumorphose/markdown"
	"github.com/skratchdot/open-golang/open"
)

type Medium struct {
	Client *medium.Medium
}

func New(token string) *Medium {
	return &Medium{Client: medium.NewClientWithAccessToken(token)}
}

func (m *Medium) Publish(md *markdown.Markdown) error {
	u, err := m.Client.GetUser("")
	if err != nil {
		return err
	}

	p, err := m.Client.CreatePost(medium.CreatePostOptions{
		UserID:        u.ID,
		Title:         md.Title,
		ContentFormat: medium.ContentFormatMarkdown,
		Content:       md.Content,
		PublishStatus: medium.PublishStatusDraft,
	})
	if err != nil {
		return err
	}

	open.Run(p.URL)
	return nil
}
