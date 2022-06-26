package main

import (
	"fmt"
	"time"

	"github.com/evertontomalok/go-concurrency/pkg/utils"
)

var c utils.Counter = utils.Counter{}

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	mainChannel := multi(ch1, ch2)

	produceMessage(ch1, ch2)

	for msg := range mainChannel {
		fmt.Println(msg)

		if c.Value() > 20 {
			close(ch1)
			close(ch2)
			break
		}
	}
}

func multi(ch1 <-chan string, ch2 <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
			case msg := <-ch1:
				out <- msg
			case msg := <-ch2:
				out <- msg
			}
		}
	}()

	return out
}

func produceMessage(ch1, ch2 chan<- string) {
	go func() {
		for {
			c.Increase()
			time.Sleep(time.Duration(500) * time.Millisecond)
			ch1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			c.Increase()
			time.Sleep(time.Duration(1500) * time.Millisecond)
			ch2 <- "Channel 2"
		}
	}()
}
