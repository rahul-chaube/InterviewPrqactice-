package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.M) {
	var myFlag string

	flag.StringVar(&myFlag, "s", "/", "Description ")

	exitCode := t.Run()
	fmt.Println("exit code is ", exitCode)
	os.Exit(exitCode)
}
