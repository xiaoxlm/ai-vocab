package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/xiaoxlm/ai-vocab/pkg/llm"
)

func TestLambdaNode(t *testing.T) {
	ctx := context.Background()
	arkLLM := llm.NewArk(os.Getenv("ARK_API_KEY"), os.Getenv("ARK_LLM_MODEL"))
	// model
	model, err := arkLLM.GetChatModel(ctx)
	if err != nil {
		t.Fatal(err)
	}
	//lambda
	lambda := compose.InvokableLambda(func(ctx context.Context, input string) (output []*schema.Message, err error) {
		desu := input + "回答结尾加上desu"
		output = []*schema.Message{
			{
				Role:    schema.User,
				Content: desu,
			},
		}
		return output, nil
	})

	chain := compose.NewChain[string, *schema.Message]()
	chain.AppendLambda(lambda).AppendChatModel(model)

	run, err := chain.Compile(ctx)
	if err != nil {
		t.Fatal(err)
	}

	answer, err := run.Invoke(ctx, "你好, 你叫什么名字")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(answer.Content)
}
