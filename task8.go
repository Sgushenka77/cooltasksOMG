package main

import (
	"fmt"
	"sync"
)

func server(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Сервер %d обработал: %s\n", id, job)
	}
}

func main() {
	serverChans := make([]chan string, 3)
	for i := range serverChans {
		serverChans[i] = make(chan string)
	}

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go server(i+1, serverChans[i], &wg)
	}

	tasks := []string{"task1", "task2", "task3", "task4", "task5", "task6", "task7"}

	for i, t := range tasks {
		idx := i % 3 // 0, 1, 2, 0, 1, 2, ...
		serverChans[idx] <- t
	}

	for i := range serverChans {
		close(serverChans[i])
	}

	wg.Wait()
	fmt.Println("Все задачи распределены")
}