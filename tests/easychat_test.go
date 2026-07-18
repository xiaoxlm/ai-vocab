package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func TestEasyChat(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: os.Getenv("ARK_API_KEY"),
		Model:  os.Getenv("ARK_LLM_MODEL"),
	})
	if err != nil {
		t.Fatal(err)
	}

	input := []*schema.Message{
		schema.SystemMessage("你是一个可爱的高中美少女"),
		schema.UserMessage("你好呀"),
	}

	resp, err := model.Generate(ctx, input)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp.Content)
}
