package tests

import (
	"context"
	"fmt"
	"github.com/xiaoxlm/ai-vocab/pkg/splitter"
	"testing"
)

func TestMarkdownTransformer(t *testing.T) {
	ctx := context.Background()

	mk := splitter.NewMarkdown(map[string]string{
		"#":   "h1",
		"##":  "h2",
		"###": "h3",
	}, true)

	splitDocs, err := mk.Transform(ctx, "../testdatas/word/A.md")
	if err != nil {
		t.Fatalf("转换失败: %v", err)
	}

	for i, doc := range splitDocs {
		fmt.Printf("片段:%d: cotent:%s\n", i+1, doc.Content)
		for k, v := range doc.MetaData {
			if k == "h1" || k == "h2" || k == "h3" {
				fmt.Printf("元数据: %s: %v\n", k, v)
			}
		}
		fmt.Println("----")
	}
}
