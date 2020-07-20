package main
import (
	"fmt"
	"time"
)
func print(id int,y *int){
	for i:=0 ; i<5 ; i++{
		fmt.Println("Goroutine", id,":", *y)
		*y += 1
	}
}
func main(){
	// Shared variable
	x := 0
	// Race condition occurs due to the outcome of program
	// which depends on interleavings

	// Race condition occurs between these two goroutines
	// The outcome (output) depends
	// on the non-deterministic order of execution
	// Switching from one goroutine to another may cause:
	//   1. Printing an older value of the shared variable
	//   2. Printing +2 value due to context switches before printing
	go print(1,&x)
	go print(2,&x)
	
	time.Sleep(time.Millisecond * 10000)
}