package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const projectId = "arbitd-182901"
const subName = "error-test-sub"

func main() {
	var mu sync.Mutex
	received := 0
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	sub := client.Subscription(subName)
	cctx, cancel := context.WithCancel(ctx)

	var records = make(map[string]int64)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		s := string(msg.Data)
		mu.Lock()
		defer mu.Unlock()
		received++
fmt.Printf("recved: %s\n", s)
		a := strings.Split(s, ",")

		n := a[0]
		publishedAt, _ := strconv.Atoi(a[1])
		records[n] = time.Now().Unix() - int64(publishedAt)

		if n == "3000" {
			cancel()
		}
		msg.Nack()
	})
	if err != nil {
		fmt.Print(err)
	}

	for n, elapsed := range records {
		fmt.Printf("%d,%d\n", n, elapsed)
	}
}
