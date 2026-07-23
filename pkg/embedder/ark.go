package embedder

import (
	"context"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

type Ark struct {
	apiKey    string
	modelName string
}

func NewArk(apiKey, modelName string) *Ark {
	return &Ark{
		apiKey:    apiKey,
		modelName: modelName,
	}
}

func (a *Ark) GetEmbedder(ctx context.Context) (*ark.Embedder, error) {
	return ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIType:               new(ark.APIType), // 为了解决 CreateEmbeddings 中 fullURL suffix 的bug问题
		MaxConcurrentRequests: new(1),           // 为了解决 CreateEmbeddings 中 fullURL suffix 的bug问题
		APIKey:                a.apiKey,
		Model:                 a.modelName,
	})
}
