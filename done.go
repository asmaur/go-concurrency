package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Working")
		}
	}
}

func main() {
	done := make(chan bool)
	go doWork(done)

	time.Sleep(time.Second * 4)
	close(done)

}
