package prodcons

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func DemoProdCons() {
	ch := make(chan int)
	wg.Add(2)
	go producer(ch)
	go consumer(ch)

	wg.Wait()
	fmt.Println("Completed ")

}

func producer(rec <-chan int) {
	defer wg.Done()
	for data := range rec {
		fmt.Println("Data received ", data)
	}
}

func consumer(send chan<- int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		send <- i
	}
	close(send)
}
