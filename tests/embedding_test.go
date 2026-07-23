package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/xiaoxlm/ai-vocab/pkg/embedder"
	"os"
	"testing"

	"github.com/joho/godotenv"
	arkopenai "github.com/sashabaranov/go-openai"
)

func TestEmbedding(t *testing.T) {
	a := embedder.NewArk(os.Getenv("ARK_API_KEY"), os.Getenv("ARK_EMBEDDING_MODEL"))

	ctx := context.Background()
	arkEmbedder, err := a.GetEmbedder(ctx)
	if err != nil {
		t.Fatal(err)
	}

	input := []string{
		"你好",
		"战神",
	}

	embeddings, err := arkEmbedder.EmbedStrings(ctx, input)
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
