package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//  In Go, the sync/atomic package provides low-level atomic memory primitives useful
// for managing shared variables among multiple goroutines in a concurrent program.
//These atomic operations are essential for ensuring data consistency and preventing race conditions without using more
// heavyweight synchronization mechanisms like mutexes.

//Atomic Load and Store: func main() {
//     var value int32
//     atomic.StoreInt32(&value, 42)
//     loadedValue := atomic.LoadInt32(&value)
//     fmt.Println(loadedValue) // prints: 42
// }
//Atomic Add: Add: Atomically adds a value to a variable and returns the new value
//Atomic Compare and Swap (CAS):
//Atomic Swap:

func atomicExample() {

	var ops atomic.Int32
	var count int32

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
				// atomic.AddInt32(&count, 1)

				count++
			}
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println("Atomic ", ops, " Normal Variable is ", count)

}

var wg sync.WaitGroup

func count(id int) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Println(id, " is executing", i)
		i := 0
		for ; i < 1e6; i++ {
		}
		_ = i
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	atomicExample()

	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(3))
	// fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(2))

	wg.Add(3)
	go count(1)
	go count(2)
	go count(3)
	wg.Wait()

	maxSum([]int{2, 3, 5, 6, 3, 6, 4, 4, 5, 9}, 3)

}

func maxSum(arr []int, k int) int {
	n := len(arr)
	if n < k {
		return 0
	}

	maxSum, windowSum := 0, 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}
	maxSum = windowSum

	for i := k; i < n; i++ {
		fmt.Println("  ***  ", i, i-k)
		windowSum = windowSum + arr[i] - arr[i-k]
		fmt.Println(maxSum, " ", windowSum)
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	fmt.Println(" Max Sum is ", maxSum)
	return maxSum
}
