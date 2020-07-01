package main
import "fmt"

func incFunc(x int) int {
	return x+1
}
func main(){
	// Declare variable as function
	var funcVar func(int) int
	funcVar = incFunc
	fmt.Println(funcVar(1))
}