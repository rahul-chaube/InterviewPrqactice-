package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "omit trail new string")
var sep = flag.String("s", "", "Sepator")

func main() {

	flag.Args()
	fmt.Print(" ", flag.Args(), *sep, *n)

	if !*n {
		fmt.Println()
	}

}
