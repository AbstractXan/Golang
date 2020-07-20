package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
var x string
var on sync.Once
func setup(){
	x = "hello"
	fmt.Println("Setting up...")
}

func doStuff(){
	on.Do(setup) // Runs only once
	fmt.Println(x)
	wg.Done()
}

func main(){
	wg.Add(2)
	go doStuff()
	go doStuff()
	wg.Wait()
}