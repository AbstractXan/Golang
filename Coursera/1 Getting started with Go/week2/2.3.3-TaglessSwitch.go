package main
import "fmt"
func main(){
	x := 5
	switch{
	case x > 1:
		fmt.Println("case 1")
	case x < -1:
		fmt.Println("case 2")
	default:
		fmt.Println("no case") 
	}
}
