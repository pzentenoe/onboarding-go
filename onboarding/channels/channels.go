package main

import (
	"fmt"
	"time"
)

//Buffered Channel
func main1() {
	intChannel := make(chan int, 2)
	intChannel <- 1
	fmt.Println(<-intChannel)

}

//UnBuffered Channel

func main() {
	intChannel := make(chan int)
	isFinished := make(chan bool, 1)

	go sendNumber(intChannel)

	go receiveNumber(intChannel, isFinished)

	fmt.Println(<-isFinished)

}

func sendNumber(intChannel chan int) {
	time.Sleep(time.Second * 1)
	for i := 0; i < 100; i++ {
		intChannel <- i + 1
	}
	close(intChannel)
}

func receiveNumber(intChannel chan int, isFinished chan bool) {
	for intChann := range intChannel {
		fmt.Println(intChann)
	}
	isFinished <- true
}
