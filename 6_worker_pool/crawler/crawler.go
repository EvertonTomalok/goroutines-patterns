package crawler

import (
	"errors"
	"time"

	"github.com/evertontomalok/go-concurrency/pkg/utils"
)

func FakeCrawl(url string) (time.Duration, error) {
	start := time.Now()
	utils.RandomSleep(50, 1000, time.Millisecond)
	took := time.Since(start)
	if took > time.Duration(700)*time.Millisecond {
		return took, errors.New("Error crawling")
	}
	return took, nil
}
