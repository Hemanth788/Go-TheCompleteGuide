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
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)

	doneChan <- true
	close(doneChan) // If you know which routine is going to be the last one, use this
}

func main() {
	// dones := make([]chan bool, 4)

	// for i := 0; i < 4; i++ {
	// 	dones[i] = make(chan bool)
	// }
	// go greet("Nice to meet you!", dones[0]) // --> runs, execution completes, nothing in the console, without the channels above that is
	// go greet("How are you?", dones[1])
	// go slowGreet("How... are... you...?", dones[2])
	// go greet("I hope you gfy!", dones[3])
	
	// for _, done := range dones {
		// 	<- done
		// }
		
		done := make(chan bool)
		go greet("Nice to meet you!", done) // --> runs, execution completes, nothing in the console, without the channels above that is
		go greet("How are you?", done)
		go slowGreet("How... are... you...?", done)
		go greet("I hope you gfy!", done)

		for doneChan := range done {
			fmt.Println(doneChan)
		}
		
}