package main

import "fmt"

// Call by value
// Good : Encapsulation
// Bad : Copy time

func foo(y int){
	y = y+2
}

func main(){
	x := 2
	foo(x)
	fmt.Println(x)
}
