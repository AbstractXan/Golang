package main

import "fmt"

func foo2(x int) (int, int){
	return x , x+1
}
func main(){
	a,b := foo2(2)
	fmt.Println(a , b)
}
