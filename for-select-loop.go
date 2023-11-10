package main

import (
	"fmt"
	"time"
)

func main() {
	// charChannel := make(chan string, 3)
	// chars := []string{"a", "ab", "acb"}

	// for _, s := range chars {
	// 	select {
	// 	case charChannel <- s:
	// 	}
	// }

	// close(charChannel)

	// for result := range charChannel {
	// 	fmt.Println(result)
	// }

	go func() {
		for {
			select {
			default:
				fmt.Println("Working")
			}
		}
	}()

	time.Sleep(time.Second * 4)

}
