package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// from https://www.youtube.com/watch?v=LSzR0VEraWw

// func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
// 	time.Sleep(d)
// 	fmt.Println(msg)
// }

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <- ctx.Done():
		log.Print(ctx.Err())
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	cancel()

	// time.AfterFunc(time.Second, cancel)
	
	// ctx, cancel := context.WithCancel(ctx)
	// go func() {
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()

	sleepAndTalk(ctx, 5 * time.Second, "hello")
}
