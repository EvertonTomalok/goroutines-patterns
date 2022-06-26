package main

import (
	"fmt"

	"github.com/evertontomalok/go-concurrency/6_worker_pool/workers"
)

const numTasks int = 50

func main() {

	tasks := make(chan string, numTasks)
	results := make(chan string, numTasks)

	workers.CreateWorkers(tasks, results, 5)

	urls := createUrlsSample()
	populateTasks(tasks, urls...)

	listenResults(results)
}

func createUrlsSample() []string {
	urls := make([]string, 0)
	for i := 0; i < numTasks; i++ {
		urls = append(urls, fmt.Sprintf("htttp://domain-%d.com.br", i))
	}

	return urls
}

func populateTasks(tasks chan string, urls ...string) {
	for _, url := range urls {
		tasks <- url
	}
	close(tasks)
}

func listenResults(results chan string) {
	for m := 0; m < numTasks; m++ {
		fmt.Println(<-results)
	}
}
