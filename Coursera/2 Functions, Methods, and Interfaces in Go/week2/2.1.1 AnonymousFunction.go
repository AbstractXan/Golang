package main
import "fmt"

func applyIt( afx func (int) int,
        val int) int{
    return afx(val) 
}
func main(){
	v := applyIt(
		func (x int) int {return x+1} ,2)
	fmt.Println(v)
}