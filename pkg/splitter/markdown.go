package splitter

import (
	"context"
	"github.com/google/uuid"
	"os"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/schema"
)

type Markdown struct {
	config *markdown.HeaderConfig
}

func NewMarkdown(headers map[string]string, trimHeaders bool) *Markdown {
	return &Markdown{
		config: &markdown.HeaderConfig{
			Headers:     headers,
			TrimHeaders: trimHeaders,
		},
	}
}

func (m *Markdown) GetSplitter(ctx context.Context) (document.Transformer, error) {
	return markdown.NewHeaderSplitter(ctx, m.config)
}

func (m *Markdown) Transform(ctx context.Context, filepath ...string) ([]*schema.Document, error) {
	var docs []*schema.Document
	for _, path := range filepath {
		fileContent, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		adoc := &schema.Document{
			ID:      uuid.New().String(),
			Content: string(fileContent),
			MetaData: map[string]any{
				"title": "A",
			},
		}
		docs = append(docs, adoc)
	}

	splitter, err := m.GetSplitter(ctx)
	if err != nil {
		return nil, err
	}

	return splitter.Transform(ctx, docs)
}
