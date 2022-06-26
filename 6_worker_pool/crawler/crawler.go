package crawler

import (
	"errors"
	"time"

	"github.com/evertontomalok/go-concurrency/pkg/utils"
)

func FakeCrawl(url string) (time.Duration, error) {
	start := time.Now()
	utils.RandomSleep(50, 500, time.Millisecond)
	took := time.Since(start)
	if took > time.Duration(400)*time.Millisecond {
		return took, errors.New("Error crawling")
	}
	return took, nil
}
