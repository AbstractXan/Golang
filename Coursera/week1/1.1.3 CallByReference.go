package main

import "fmt"

// Call by Reference
// Good : Saves copying time. No copy required.
// Bad : Data Encapsulation not supported.

func foo(y *int){
	*y = *y+2
}

func main(){
	x := 2
	foo(&x)
	fmt.Println(x)
}
