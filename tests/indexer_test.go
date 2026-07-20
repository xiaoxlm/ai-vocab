package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	qdrant_indexer "github.com/cloudwego/eino-ext/components/indexer/qdrant"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/qdrant/go-client/qdrant"
)

func TestIndexer(t *testing.T) {
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

	// indexer
	indexer, err := qdrant_indexer.NewIndexer(ctx, &qdrant_indexer.Config{
		Client:     qdrantClient,
		Collection: "test",
		VectorDim:  2048,
		Distance:   qdrant.Distance_Cosine,
		Embedding:  embedder, // 你的 embedding 组件
	})
	if err != nil {
		t.Fatal(err)
	}

	// storage
	docs := []*schema.Document{
		{
			ID:      uuid.New().String(),
			Content: "这个世界很操蛋，因为太多坏人和自私的人了。无视法律与人个体的祖研",
			MetaData: map[string]any{
				"title":  "这个世界很操蛋",
				"chunk":  1,
				"author": "Oscar",
			},
		},
		{
			ID:      uuid.New().String(),
			Content: "如果这个世界多一些善良的人的话,一切将会更美好",
			MetaData: map[string]any{
				"title":  "这个世界很操蛋",
				"chunk":  2,
				"author": "Oscar",
			},
		},
	}
	ids, err := indexer.Store(ctx, docs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ids:", ids)

}
