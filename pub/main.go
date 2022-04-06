package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
)

const projectId = "arbitd-182901"
const topicId = "error-test"

func main() {
	// 空のコンテキスト作成
	ctx := context.Background()
	// クライアント作成
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		fmt.Print("client error.err:", err)
		os.Exit(1)
	}
	// トピックへの参照を作成
	t := client.Topic(topicId)

	if len(os.Args) < 2  {
		fmt.Println("missing argument")
		os.Exit(2)
	}

	n := os.Args[1]

	// メッセージを発行
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(fmt.Sprintf("%s", n)),
	})
	// メッセージIDとPublish呼び出しのエラーを発行
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
}
