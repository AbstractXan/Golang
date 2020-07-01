package main
import (
	"fmt"
	"math"
)

// Returns a distance function with a set origin
func MakeDistOrigin(originX, originY float64) func (float64, float64) float64 { 
	fn := func (x , y float64) float64 { 
		return math.Sqrt( math.Pow(x-originX,2) + math.Pow(y-originY,2) )
	}
	return fn
}
func main(){
	Dist1 := MakeDistOrigin(0,0)
	Dist2 := MakeDistOrigin(1,1)
	fmt.Println("(1,1) from (0,0) : ",Dist1(1,1))
	fmt.Println("(1,1) from (1,1) : ",Dist2(1,1))
}