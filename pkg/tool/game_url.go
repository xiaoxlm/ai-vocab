package tool

import (
	"context"
	"fmt"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

type GameURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type InputParams struct {
	Name string `json:"name" jsonschema:"type=string,description=the name of the game"`
}

func GetGameURL(_ context.Context, params *InputParams) (string, error) {
	games := []GameURL{
		{Name: "原神", URL: "https://ys.mihayo.com/tool"},
		{Name: "英雄联盟", URL: "https://lol.qq.com/tool"},
		{Name: "DOTA2", URL: "https://www.dota2.com/tool"},
	}

	for _, game := range games {
		if game.Name == params.Name {
			return game.URL, nil
		}
	}

	return "", fmt.Errorf("game not found")
}

func CreateGameURLTool() tool.InvokableTool {
	return utils.NewTool(
		&schema.ToolInfo{
			Name: "GetGameURL",
			Desc: "get a game's url by game name",
			ParamsOneOf: schema.NewParamsOneOfByParams(
				map[string]*schema.ParameterInfo{
					"name": {
						Type:     schema.String,
						Desc:     "the name of the game",
						Required: true,
					},
				},
			),
		},
		GetGameURL)
}
