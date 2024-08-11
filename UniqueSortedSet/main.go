package main

import (
	"fmt"
	"time"
	sortedset "uniquesortedset/sortedSet"
)

func main() {
	set := sortedset.NewQue()
	set.Add("Item1")
	time.Sleep(time.Second)
	set.Add("Item2")
	time.Sleep(time.Second)
	set.Add("Item3")

	fmt.Printf(" Print odest %+v \n", set.GetOlder())
	fmt.Printf(" Print odest %+v \n", set.GetOlder())
	fmt.Printf(" Print odest %+v \n", set.GetOlder())
	fmt.Printf(" Print odest %+v \n", set.GetOlder())

	maps := make(map[Student]bool)
	St1 := Student{Name: "Rahul"}

	maps[St1] = true

	St2 := Student{Name: "Rahul"}

	maps[St2] = true

	fmt.Println(" ", len(maps), maps)

}

type Student struct {
	Name string
	Id   int
}
