package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // only works if you know which operation takes the longes
}

func main() {
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)

	go slowGreet("How ... are ... you ...?", done)

	go greet("I hope you're liking the course!", done)
	for doneChan := range done {
		fmt.Println(doneChan)
	}
	// <-done //will contine after done us true
}
