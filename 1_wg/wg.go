package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/evertontomalok/go-concurrency/pkg/utils"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go RandomSleepAndPrint(i, &wg)
	}
	wg.Wait()

	fmt.Println("You're so lazy!")
}

func RandomSleepAndPrint(sequence int, wg *sync.WaitGroup) {
	defer wg.Done()
	sleepFor := utils.RandInt(100, 500)
	time.Sleep(time.Duration(sleepFor) * time.Microsecond)
	fmt.Printf("%d Sleeping for %d microseconds...\n", sequence, sleepFor)
}
