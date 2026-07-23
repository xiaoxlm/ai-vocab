package retriever

import (
	"context"

	"github.com/cloudwego/eino-ext/components/indexer/qdrant"
    qdrant_retriever "github.com/cloudwego/eino-ext/components/retriever/qdrant"
)

type Qdrant struct {
	Host string
	Port int
}

func NewQdrant(host string, port int) *Qdrant {
	return &Qdrant{
		Host: host,
		Port: port,
	}
}

func (q *Qdrant) GetRetriever(ctx context.Context, ) (*qdrant_retriever.Retriever, error) {
	return qdrant_retriever.NewRetriever(ctx, &qdrant.Config{
		Host: q.Host,
		Port: q.Port,
	})
}
