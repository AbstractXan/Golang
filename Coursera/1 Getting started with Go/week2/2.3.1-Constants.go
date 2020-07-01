package main
import "fmt"
type Grades int
const (
	A Grades = iota
	B
	C
	D
	F
)
func main(){
	fmt.Println(A,B,C,D,F)
}


