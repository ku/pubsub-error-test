package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
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

	if len(os.Args) < 2 {
		fmt.Println("missing argument")
		os.Exit(2)
	}

	lim, _ := strconv.Atoi(os.Args[1])

	for i := 0; i <= lim; i++ {
		println(i)
		time.Sleep(1000 * time.Millisecond)
		result := t.Publish(ctx, &pubsub.Message{
			Data: []byte(fmt.Sprintf("%d,%d", i, time.Now().Unix())),
		})
		_, err := result.Get(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}
}
