package main

import "fmt"

func foo(x int) int {
	return x + 1
}

func main(){
	y := foo(1)
	fmt.Printf("%d\n",y)
}
