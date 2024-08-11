package main

import (
	"fmt"
)

func main() {
	// fmt.Println(" ********  Design Pattern Example ********  ")

	// fmt.Println(" Singletone Demo ")

	// signletone.DemoSingletone()

	// fmt.Println(" Factory method ")

	// factory.DemoFactory()

	Hi("Rahul Chuube ")
	v1 := new(int)
	var v2 *int
	var in interface{}

	Hi(v1)
	Hi(v2)
	Hi(in)

}

func Hi(in interface{}) {

	fmt.Println(" ", in, in == nil)

	switch in.(type) {
	case nil:
		fmt.Println(" Value is nil ")
	case interface{}:
		{
			fmt.Println(" Interface ")
		}
	}

}
