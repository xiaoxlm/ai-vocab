package indexer

import (
	"context"

	qdrant_indexer "github.com/cloudwego/eino-ext/components/indexer/qdrant"
	qdrant_retriever "github.com/cloudwego/eino-ext/components/retriever/qdrant"
	"github.com/cloudwego/eino/components/embedding"
	"github.com/qdrant/go-client/qdrant"
)

type Qdrant struct {
	client   *qdrant.Client
	embedder embedding.Embedder
}

func NewQdrant(client *qdrant.Client, embedder embedding.Embedder) *Qdrant {
	return &Qdrant{
		client:   client,
		embedder: embedder,
	}
}

func (q *Qdrant) GetIndexer(ctx context.Context, config *qdrant_indexer.Config) (*qdrant_indexer.Indexer, error) {
	if config.Client == nil {
		config.Client = q.client
	}
	if config.Embedding != nil {
		config.Embedding = q.embedder
	}

	return qdrant_indexer.NewIndexer(ctx, config)
}

func (q *Qdrant) GetRetriever(ctx context.Context, config *qdrant_retriever.Config) (*qdrant_retriever.Retriever, error) {
	if config.Client == nil {
		config.Client = q.client
	}
	if config.Embedding == nil {
		config.Embedding = q.embedder
	}
	return qdrant_retriever.NewRetriever(ctx, config)
}
