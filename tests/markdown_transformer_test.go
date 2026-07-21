package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
)

func TestMarkdownTransformer(t *testing.T) {
	ctx := context.Background()

	splitter, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":   "h1",
			"##":  "h2",
			"###": "h3",
		},
		TrimHeaders: true,
	})
	if err != nil {
		t.Fatalf("创建拆分器失败: %v\n", err)
	}

	fileContent, err := os.ReadFile("../testdatas/word/A.md")
	if err != nil {
		t.Fatalf("读取文件失败: %v\n", err)
	}
	adoc := &schema.Document{
		ID:      uuid.New().String(),
		Content: string(fileContent),
		MetaData: map[string]any{
			"title": "A",
		},
	}

	splitDocs, err := splitter.Transform(ctx, []*schema.Document{adoc})
	if err != nil {
		log.Fatalf("转换失败: %v", err)
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
