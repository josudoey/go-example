package cancel_test

import (
	"context"
	"fmt"
	"sync"
)

func ExampleContext_Cancel() {
	ctx, cancel := context.WithCancel(context.Background())
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go (func() {
		defer waitGroup.Done()
		<-ctx.Done()
		fmt.Printf("background done\n")
	})()
	cancel()
	waitGroup.Wait()
	if ctx.Err() == context.Canceled {
		fmt.Printf("canceled\n")
	}
	fmt.Printf("done\n")

	// Output:
	// background done
	// canceled
	// done

}
