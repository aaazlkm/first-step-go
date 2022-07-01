package main

import (
	"fmt"
	"time"
)

func fib(c chan int, quit chan bool) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
		fmt.Println("in for")
	}
}

func main() {
	start := time.Now()

	command := ""
	data := make(chan int)
	quit := make(chan bool)

	go fib(data, quit)

	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	time.Sleep(1 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
