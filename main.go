package main

import "github.com/evertontomalok/go-concurrency/pkg/utils"

func main() {
	msg := "Doing task..."
	utils.Spinner(30, msg)
}
