package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	for { //2. 	Will cause deadlock if channel doesn't close
		msg, open := <-c
		if !open {
			break
		}
		/*
		1. Channels are 'blocking' in nature
		Won't let main exit */
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)

	/*
		3. Always a good practice to close channel from sender's end.
		Closing a channel from receivers end might cause sender to panic.
	*/
}
