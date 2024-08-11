package main

import (
	pubsub "GoRoutineDesignPattern/PubSub"
	"GoRoutineDesignPattern/pipeline"
	"fmt"
)

func main() {
	fmt.Println("  Pub sub model ")
	pubsub.PublisherDemo()

	fmt.Println(" Piple line model ")

	pipeline.DemoPipeline()
}
