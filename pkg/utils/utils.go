package utils

import (
	"math/rand"
	"sync"
	"time"
)

type Counter struct {
	Mu    sync.Mutex
	Count int
}

func (c *Counter) Increase() {
	c.Mu.Lock()
	c.Count++
	c.Mu.Unlock()
}

func (c *Counter) Decrease() {
	c.Mu.Lock()
	c.Count--
	c.Mu.Unlock()
}

func (c *Counter) Value() int {
	c.Mu.Lock()
	var value int = c.Count
	c.Mu.Unlock()

	return value
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func RandomSleep(min int, max int, duration time.Duration) {
	time.Sleep(time.Duration(RandInt(min, max)) * duration)
}
