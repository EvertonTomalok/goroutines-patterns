package main

import (
	"fmt"

	"github.com/evertontomalok/go-concurrency/6_worker_pool/workers"
	"github.com/evertontomalok/go-concurrency/pkg/utils"
)

const numWorkers int = 20

func main() {
	urls := createUrlsSample()
	var numTasks int = len(urls)

	tasks := make(chan string)
	results := make(chan string)

	workers.CreateWorkers(tasks, results, numWorkers)

	go populateTasks(tasks, urls...)

	listenResults(results, numTasks)
}

func createUrlsSample() []string {
	// simulating crawling urls
	msg := "Collecting URLS"
	utils.Spinner(15, msg)

	totalUrls := utils.RandInt(400, 500)

	urls := make([]string, totalUrls)
	for i := 0; i < totalUrls; i++ {
		urls[i] = fmt.Sprintf("htttp://domain-%d.com.br", i)
	}

	return urls
}

func populateTasks(tasks chan string, urls ...string) {
	for _, url := range urls {
		tasks <- url
	}
	close(tasks)
}

func listenResults(results chan string, totalResults int) {
	for m := 0; m < totalResults; m++ {
		fmt.Println(<-results)
	}
}
