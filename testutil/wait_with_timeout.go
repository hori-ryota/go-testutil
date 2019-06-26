package testutil

import (
	"context"
	"sync"
	"time"
)

func WaitWithTimeout(wg *sync.WaitGroup, timeout time.Duration) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return WaitWithContext(ctx, wg)
}

func WaitWithContext(ctx context.Context, wg *sync.WaitGroup) error {
	doneChan := make(chan struct{})
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case <-doneChan:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
