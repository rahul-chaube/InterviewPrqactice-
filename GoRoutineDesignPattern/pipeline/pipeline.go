package pipeline

import "fmt"

func genrateNumber() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch

}

func square(in <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for data := range in {
			out <- data * data
		}
		close(out)

	}()

	return out
}

func DemoPipeline() {
	number := genrateNumber()

	out := square(number)

	for data := range out {
		fmt.Println("  ", data)
	}

	fmt.Println("Exit ")
}
