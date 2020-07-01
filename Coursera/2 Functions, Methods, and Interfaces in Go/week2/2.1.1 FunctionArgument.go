package main
import "fmt"

func applyIt( afx func (int) int,
        val int) int{
    return afx(val) 
}
func incFunc(x int) int {
	return x+1
}
func main(){
	v := applyIt(incFunc,2)
	fmt.Println(v)
}