package main

import (
	"fmt"
	"time"

	"github.com/evertontomalok/go-concurrency/pkg/utils"
)

func main() {
	c := utils.Counter{}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Duration(500) * time.Millisecond)
			ch1 <- "Channel 1"

			c.Mu.Lock()
			c.Count++
			c.Mu.Unlock()
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(1500) * time.Millisecond)
			ch2 <- "Channel 2"

			c.Mu.Lock()
			c.Count++
			c.Mu.Unlock()
		}
	}()

	for {
		if c.Count > 20 {
			fmt.Println("Stopping")
			break
		}

		select {
		case channel1 := <-ch1:
			fmt.Println(channel1)
		case channel2 := <-ch2:
			fmt.Println(channel2)
		}
	}
}
