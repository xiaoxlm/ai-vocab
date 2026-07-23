package tests

import (
	"context"
	"fmt"
	"github.com/xiaoxlm/ai-vocab/pkg/embedder"
	"github.com/xiaoxlm/ai-vocab/pkg/indexer"
	"os"
	"testing"

	qdrant_retriever "github.com/cloudwego/eino-ext/components/retriever/qdrant"
)

func TestRetriever(t *testing.T) {
	ctx := context.Background()
	// embedder
	arkEmbedder, err := embedder.NewArk(os.Getenv("ARK_API_KEY"), os.Getenv("ARK_EMBEDDING_MODEL")).GetEmbedder(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// 创建 retriever
	retriever, err := indexer.NewQdrant(qdrantClient, arkEmbedder).GetRetriever(ctx, &qdrant_retriever.Config{
		//Client:         qdrantClient,
		Collection: "test",
		//Embedding:      embedder,
		ScoreThreshold: new(0.4),
		TopK:           5,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 搜索
	docs, err := retriever.Retrieve(ctx, "操蛋")
	if err != nil {
		t.Fatal(err)
	}
	for _, doc := range docs {
		fmt.Println("=========")
		fmt.Println(doc.ID)
		fmt.Println(doc.Content)
		fmt.Println(doc.MetaData)
	}
}
