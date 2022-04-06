package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"sync"
)

const projectId = "arbitd-182901"
const subName = "error-test-sub"

func main() {
	var mu sync.Mutex
	received := 0
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	// サブスクリプションの参照作成
	sub := client.Subscription(subName)
	// レシーブを終了するためのコンテキスト作成
	cctx, cancel := context.WithCancel(ctx)
	_ = cancel
	// メッセージ受信
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		//		msg.Ack()
		s := string(msg.Data)
		fmt.Printf("Got message: %q\n", s)
		mu.Lock()
		defer mu.Unlock()
		received++

		//n, err := strconv.Atoi(s)
		//if err != nil {
		//	fmt.Printf("failed to parse %s\n", err.Error())
		//	msg.Ack()
		//}

//		fmt.Printf("%d\n", n)
		msg.Nack()
	})
	if err != nil {
		fmt.Print(err)
	}
}
