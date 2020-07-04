package main
import "fmt"
type Shape2D interface{
	Area() float64
	Perimeter() float64
}

type Triangle {}
func (t Triangle) Area() float64 {}
func (t Triangle) Perimeter() float64 {}

func main(){

}