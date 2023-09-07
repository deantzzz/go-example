package context_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestCancel 测试context的cancel方法
func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()
	go loop(ctx, "1")
	go loop(ctx, "2")
	go loop(ctx, "3")
	for {
	}
}

func loop(ctx context.Context, i string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done-" + i)
			return
		default:
			fmt.Println("work-" + i)
			time.Sleep(time.Second)
		}
	}
}
