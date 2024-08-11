package pubsub

import (
	"fmt"
	"sync"
	"time"
)

type Punlisher struct {
	Subscriber []chan int
}

func (p *Punlisher) Subscribe() chan int {
	ch := make(chan int)
	p.Subscriber = append(p.Subscriber, ch)
	return ch
}

func (p *Punlisher) Publish(num int) {
	for i, sub := range p.Subscriber {
		fmt.Println(i, "Chanel send ", num)
		sub <- num
	}
}

func PublisherDemo() {
	pub := Punlisher{}

	ch1 := pub.Subscribe()
	ch2 := pub.Subscribe()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for data := range ch1 {
			// if !ok {
			// 	fmt.Println(" Chanel is closed ")
			// 	return
			// }
			fmt.Println("Channel 1 received ", data)
		}
	}()
	go func() {
		defer wg.Done()

		for data := range ch2 {
			// if !ok {
			// 	fmt.Println(" Chanel is closed ")
			// 	return
			// }
			fmt.Println("Channel 2 received ", data)
		}
	}()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		pub.Publish(i)
	}

	close(ch1)
	close(ch2)

	wg.Wait()

}
