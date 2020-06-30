package main
import (
"fmt"
)

type person struct{
	name string
	age int
}

func main(){
	p := person{name: "Nate", age: 14}
	fmt.Println(p)
	fmt.Println(p.name)
}
