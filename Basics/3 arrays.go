package main
import "fmt"

func main(){
	var a [5]int
	fmt.Println(a) // Zero valued

	//Short hand
	b := [5]int{1,2,3,4,5}
	fmt.Println(b)

	//Dynamic array-> Slices
	c := []int{1,2,3,4,5}
	c = append(c,13)
	fmt.Println(c)

	//Map: Key value pairs
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"]=4

	delete(vertices,"triangles")
	fmt.Println(vertices)
}
