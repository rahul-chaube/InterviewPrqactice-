package workpool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//

func MyFunc(ctx context.Context, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case data := <-ch:
			fmt.Println(" Data received  ", data)
		case <-ctx.Done():
			fmt.Println("Time out occured ")

			return

		}
	}

}

var wg sync.WaitGroup

func DemoCancelStatement() {
	ch := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	wg.Add(1)
	// go func() {
	// 	time.Sleep(5 * time.Second)
	// 	cancel()
	// 	wg.Done()
	// }()
	go MyFunc(ctx, ch)

	go func() {
		i := 0
		for {
			i++
			time.Sleep(time.Second)
			ch <- i
		}
	}()

	wg.Wait()
	time.Sleep(time.Second)
}
