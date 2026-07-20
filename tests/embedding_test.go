package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/joho/godotenv"
	arkopenai "github.com/sashabaranov/go-openai"
)

func TestEmbedding(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIType:               new(ark.APIType), // 为了解决 CreateEmbeddings 中 fullURL suffix 的bug问题
		MaxConcurrentRequests: new(1),           // 为了解决 CreateEmbeddings 中 fullURL suffix 的bug问题
		APIKey:                os.Getenv("ARK_API_KEY"),
		Model:                 "doubao-embedding-vision-251215",
	})
	if err != nil {
		t.Fatal(err)
	}

	input := []string{
		"你好",
		"战神",
	}

	embeddings, err := embedder.EmbedStrings(ctx, input)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range embeddings {
		fmt.Printf("e length:%d, e content:%v\n", len(e), e)
	}
}

func TestEmbedding2(t *testing.T) {
	err := godotenv.Load("../.env")
	ARK_API_KEY := os.Getenv("ARK_API_KEY")
	config := arkopenai.DefaultConfig(ARK_API_KEY)
	config.BaseURL = "https://ark.cn-beijing.volces.com/api/v3"
	client := arkopenai.NewClientWithConfig(config)
	fmt.Println("----- embeddings request -----")
	req := arkopenai.EmbeddingRequestStrings{
		Input: []string{
			"花椰菜又称菜花、花菜，是一种常见的蔬菜。",
		},
		Model:          "doubao-embedding-vision-251215",
		EncodingFormat: arkopenai.EmbeddingEncodingFormatFloat,
	}

	resp, err := client.CreateEmbeddings(context.Background(), req)
	if err != nil {
		fmt.Printf("embeddings error: %v\n", err)
		return
	}

	s, _ := json.Marshal(resp)
	fmt.Println(string(s))
}
