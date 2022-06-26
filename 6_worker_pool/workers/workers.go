package workers

import (
	"fmt"
	"time"

	"github.com/evertontomalok/go-concurrency/6_worker_pool/crawler"
	"github.com/evertontomalok/go-concurrency/pkg/colors"
)

func Worker(tasks <-chan string, results chan<- string) {
	for url := range tasks {
		took, err := crawler.FakeCrawl(url)
		if err != nil {
			results <- fmt.Sprintf("%s[FAILED] %s crawled. (%d ms) %s", colors.RED, url, took/time.Millisecond, colors.END)
			continue
		}
		results <- fmt.Sprintf("%s[SUCCESSFUL] %s crawled. (%d ms) %s", colors.GREEN, url, took/time.Millisecond, colors.END)
	}

}

func CreateWorkers(tasks <-chan string, results chan<- string, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go Worker(tasks, results)
	}
}
