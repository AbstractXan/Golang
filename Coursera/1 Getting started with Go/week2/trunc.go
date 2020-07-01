package main
import "fmt"
func main(){
	var f float32
	fmt.Println("Enter a float number")
	fmt.Scan(&f)
	n := int16(f)
	fmt.Println(n)
}
