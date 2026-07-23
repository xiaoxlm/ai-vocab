package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"

	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	callbackshepler "github.com/cloudwego/eino/utils/callbacks"
	"github.com/xiaoxlm/ai-vocab/pkg/llm"
	pkgtool "github.com/xiaoxlm/ai-vocab/pkg/tool"
)

func TestInvokeGameTool(t *testing.T) {
	// with handler
	modelHandler := &callbackshepler.ModelCallbackHandler{
		OnEnd: func(ctx context.Context, _ *callbacks.RunInfo, output *model.CallbackOutput) context.Context {
			fmt.Printf("the thinking process: %s\n", output.Message.Content)
			return ctx
		},
	}

	toolHandler := &callbackshepler.ToolCallbackHandler{
		OnStart: func(ctx context.Context, _ *callbacks.RunInfo, input *tool.CallbackInput) context.Context {
			fmt.Printf("begin to execute tools, the params are: %s\n", input.ArgumentsInJSON)
			return ctx
		},
		OnEnd: func(ctx context.Context, _ *callbacks.RunInfo, output *tool.CallbackOutput) context.Context {
			fmt.Printf("the tool execution result: %s\n", output.Response)
			return ctx
		},
	}

	handler := callbackshepler.NewHandlerHelper().ChatModel(modelHandler).Tool(toolHandler).Handler()

	ctx := context.Background()
	arkLLM := llm.NewArk(os.Getenv("ARK_API_KEY"), os.Getenv("ARK_LLM_MODEL"))
	// get model
	arkModel, err := arkLLM.GetChatModel(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// bindTool
	gameURLTool := pkgtool.CreateGameURLTool()
	info, err := gameURLTool.Info(ctx)
	if err != nil {
		t.Fatal(err)
	}

	toolInfos := []*schema.ToolInfo{info}
	if err = arkModel.BindTools(toolInfos); err != nil {
		t.Fatal(err)
	}

	// create chain
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	// create tool node
	toolNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: []tool.BaseTool{gameURLTool},
	})
	if err != nil {
		t.Fatal(err)
	}
	// connect and compile
	chain.AppendChatModel(arkModel, compose.WithNodeName("chat_model")).AppendToolsNode(toolNode, compose.WithNodeName("tools"))

	agent, err := chain.Compile(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// running
	resp, err := agent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "请告诉我'英雄联盟','原神','DOTA2'的URL是什么", // 目前这里还有个问题，就是如果dota2是小写，就会报错
		},
	}, compose.WithCallbacks(handler))
	if err != nil {
		t.Fatal(err)
	}

	for _, msg := range resp {
		fmt.Println(msg.Content)
	}
}
