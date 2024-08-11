package workpool

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increament() {
	defer mu.Unlock()
	mu.Lock()
	counter++

}

func IninitMutex() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); increament() }()
	}

	wg.Wait()

	fmt.Println(" Increament value ", counter)
}
