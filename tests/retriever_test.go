package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	qdrant_retriever "github.com/cloudwego/eino-ext/components/retriever/qdrant"
	"github.com/joho/godotenv"
	"github.com/qdrant/go-client/qdrant"
)

func TestRetriever(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	// embedder
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIType:               new(ark.APIType), // 为了解决 CreateEmbeddings 中 fullURL suffix 的bug问题
		MaxConcurrentRequests: new(1),           // 为了解决 CreateEmbeddings 中 fullURL suffix 的bug问题
		APIKey:                os.Getenv("ARK_API_KEY"),
		Model:                 "doubao-embedding-vision-251215",
	})
	if err != nil {
		t.Fatal(err)
	}
	// qdrant client
	qdrantClient, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		t.Fatal(err)
	}

	// 创建 retriever
	retriever, _ := qdrant_retriever.NewRetriever(ctx, &qdrant_retriever.Config{
		Client:         qdrantClient,
		Collection:     "test",
		Embedding:      embedder,
		ScoreThreshold: new(0.4),
		TopK:           5,
	})

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
