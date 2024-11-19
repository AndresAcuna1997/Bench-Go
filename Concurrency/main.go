package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	//dones := make([]chan bool, 4)
	isDone := make(chan bool)

	//dones[0] = make(chan bool)
	go greet("Nice to meet you!", isDone)
	//dones[1] = make(chan bool)
	go greet("How are you?", isDone)
	//dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", isDone)
	//dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", isDone)

	// for _, done := range dones {
	// 	<-done
	// }

	for doneChan := range isDone {
		fmt.Println(doneChan)
	}
}
