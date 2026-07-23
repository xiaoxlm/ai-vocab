package llm

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

type Ark struct {
	apiKey    string
	modelName string
}

func NewArk(apiKey, modelName string) *Ark {
	return &Ark{
		apiKey:    apiKey,
		modelName: modelName,
	}
}

func (a *Ark) GetChatModel(ctx context.Context) (*ark.ChatModel, error) {
	return ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: a.apiKey,
		Model:  a.modelName,
	})
}
