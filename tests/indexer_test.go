package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/xiaoxlm/ai-vocab/pkg/embedder"
	"github.com/xiaoxlm/ai-vocab/pkg/indexer"

	qdrant_indexer "github.com/cloudwego/eino-ext/components/indexer/qdrant"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
)

var qdrantClient *qdrant.Client

func init() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	qdrantClient = client
}

func TestIndexer(t *testing.T) {
	ctx := context.Background()
	arkEmbedder, err := embedder.NewArk(os.Getenv("ARK_API_KEY"), os.Getenv("ARK_EMBEDDING_MODEL")).GetEmbedder(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// indexer
	qdrantIndexer, err := indexer.NewQdrant(qdrantClient, arkEmbedder).GetIndexer(ctx, &qdrant_indexer.Config{
		//Client:     qdrantClient,
		Collection: "test",
		VectorDim:  2048,
		Distance:   qdrant.Distance_Cosine,
		//Embedding:  arkEmbedder, // 你的 embedding 组件
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
	ids, err := qdrantIndexer.Store(ctx, docs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ids:", ids)

}
