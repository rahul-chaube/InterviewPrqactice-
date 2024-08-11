package main

import (
	prodcons "GoRoutineDemo/prodCons"
	"GoRoutineDemo/workpool"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	CreateWorker(5)
	wg.Wait()

	fmt.Println("Exit ****")

	workpool.IninitMutex()

	workpool.OnceInit()

	// workpool.Pool.Put(" 1 ")
	// workpool.Pool.Put(" 2 ")
	// workpool.Pool.Put(" 3 ")

	// workpool.Pool.Get()
	// workpool.Pool.Get()
	// workpool.Pool.Get()

	workpool.DemoCancelStatement()
	fmt.Println(" Producer Consumer problem ")
	prodcons.DemoProdCons()

}

func Worker(ch1 chan int, workerName string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		job, ok := <-ch1
		if !ok {
			return
		}
		time.Sleep(time.Second)
		fmt.Printf("%d is completed by %s at %s\n", job, workerName, time.Now().Format(time.RFC3339))
	}
}

func CreateWorker(numberOfWorker int) {
	ch := make(chan int, 10)
	wg.Add(numberOfWorker) // Add the number of workers to the WaitGroup
	for i := 0; i < numberOfWorker; i++ {
		go Worker(ch, fmt.Sprintf("Worker-%d", i), &wg)
	}

	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Sending work ", i)
	}
	close(ch)
	fmt.Println("Work is completed")
	wg.Done() // Mark CreateWorker as done
}
