package main

import "fmt"

type Command struct {
	op    string // "add", "get"
	value int
	reply chan int
}

func manager(commands chan Command) {
	state := 0 // только эта горутина трогает state

	for cmd := range commands {
		switch cmd.op {
		case "add":
			state = state + cmd.value
			cmd.reply <- state
		case "get":
			cmd.reply <- state
		}
	}
}

func main() {
	commands := make(chan Command)
	go manager(commands)

	reply1 := make(chan int)
	commands <- Command{"add", 10, reply1}
	fmt.Println("После add 10:", <-reply1)

	reply2 := make(chan int)
	commands <- Command{"add", 5, reply2}
	fmt.Println("После add 5:", <-reply2)

	reply3 := make(chan int)
	commands <- Command{"get", 0, reply3}
	fmt.Println("Текущее:", <-reply3)
}