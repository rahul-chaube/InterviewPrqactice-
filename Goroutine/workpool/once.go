package workpool

import (
	"fmt"
	"sync"
)

var once sync.Once

func initMethod() {
	fmt.Println("Once is initialized ")
}

func OnceInit() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Calling ")
			defer wg.Done()
			once.Do(initMethod)
		}()
	}

	wg.Wait()
}
