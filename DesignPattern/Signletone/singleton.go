package signletone

import (
	"fmt"
	"sync"
)

//Single Instance: Ensures that a class has only one instance.
// Global Access Point: Provides a global access point to that instance.

type Singleton struct {
	Value int
}

var instance *Singleton

var once sync.Once

func GetInstance() *Singleton {

	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func DemoSingletone() {
	s1 := GetInstance()
	s1.Value = 43

	fmt.Println(" Value is ", s1.Value)

	s2 := GetInstance()

	fmt.Println(" Second velue   ", s2.Value)

	if s1 == s2 {
		fmt.Println(" Same same ")
	}
}
