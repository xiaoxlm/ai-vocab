package tests

import (
	"context"
	"fmt"
	"github.com/xiaoxlm/ai-vocab/pkg/llm"
	"os"
	"testing"

	"github.com/cloudwego/eino/schema"
	_ "github.com/xiaoxlm/ai-vocab/config"
)

func TestEasyChat(t *testing.T) {
	ctx := context.Background()
	arkLLM := llm.NewArk(os.Getenv("ARK_API_KEY"), os.Getenv("ARK_LLM_MODEL"))
	model, err := arkLLM.GetChatModel(ctx)
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
