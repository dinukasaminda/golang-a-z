package main

import (
	"context"
	"fmt"
	"time"
)

func query(ctx context.Context) error {
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("query complete")
		return nil
	case <-ctx.Done():
		// we can handle any canceltion or termination of this task
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)

	defer cancel()
	fmt.Println(query(ctx))
}
