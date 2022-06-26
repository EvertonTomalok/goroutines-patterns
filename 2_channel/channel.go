package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/evertontomalok/go-concurrency/pkg/utils"
)

type loopController struct {
	mu           sync.Mutex
	runningTasks int
}

func (l *loopController) increment() {
	l.mu.Lock()
	l.runningTasks++
	l.mu.Unlock()
}

func (l *loopController) decrement() {
	l.mu.Lock()
	if l.runningTasks > 0 {
		l.runningTasks--
	}
	l.mu.Unlock()
}

func main() {
	p := producer()

	for st := range p {
		fmt.Println(st)
	}

}

func producer() <-chan string {
	c := make(chan string)

	controller := loopController{}
	for i := 0; i < 10; i++ {
		controller.increment()
		go sleepAsync(c, i, &controller)
	}
	return c
}

func sleepAsync(channel chan<- string, sequence int, controller *loopController) {
	sleepFor := utils.RandInt(100, 1500)
	channel <- fmt.Sprintf("%d Sleeping for %d Millisecond...\n", sequence, sleepFor)
	time.Sleep(time.Duration(sleepFor) * time.Millisecond)
	channel <- fmt.Sprintf("%d Done\n", sequence)

	controller.decrement()

	if controller.runningTasks == 0 {
		close(channel)
	}
}
