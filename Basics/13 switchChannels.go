package main
import (
"fmt"
"time"
)
func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go every500ms(c1)
	go every2sec(c2)

	for {
		select {
			case msg1 := <- c1:
				fmt.Println(msg1)
			case msg2 := <- c2:
				fmt.Println(msg2)
		}
	}
}
func  every500ms(c chan string){
	for{
		c <- "Every 500 ms"
		time.Sleep(time.Millisecond * 500)
	}
}

func every2sec(c chan string){
	for{
		c <- "Every two seconds"
		time.Sleep(time.Second * 2)
	}
}
