package main

import "fmt"

// Slices contain a pointer to the array
// Passing slices copies a pointer

func foo(sli int){
	sli[0] = sli[0] + 1 
}

func main(){
	a := []int{1,2,3}
	foo(a)
	fmt.Println(a)
}
