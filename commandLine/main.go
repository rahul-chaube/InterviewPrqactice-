package main

import (
	"flag"
	"fmt"
)

func main() {
	nameFlag := flag.String("name", "", "Please Specified name ")
	ageFlag := flag.Int("age", 0, "Please enter user name ")
	boolFlag := flag.Bool("debug", false, "Define mode of debuging")

	flag.Usage = func() {
		fmt.Println("This program is basic exaple of Command line [option]")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println(" Enter name is  ", *nameFlag, " age is ", *ageFlag, " is Debug mode on ", *boolFlag)

}
