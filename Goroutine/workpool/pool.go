package workpool

import (
	"fmt"
	"sync"
)

var Pool sync.Pool

func Put(data string) {
	Pool.Put(func(input string) {
		fmt.Println(" Out Operation ", data)
	})
}
