package main

import (
	"fmt"
	"time"
)

func request(source string, delay time.Duration) string {
	time.Sleep(delay)
	return source + " -> данные"
}

func main() {
	responses := make(chan string, 3) 

	go func() { responses <- request("База 1", 300*time.Millisecond) }()
	go func() { responses <- request("База 2", 100*time.Millisecond) }()
	go func() { responses <- request("База 3", 500*time.Millisecond) }()

	first := <-responses 
	fmt.Println("Первый ответ:", first)
}