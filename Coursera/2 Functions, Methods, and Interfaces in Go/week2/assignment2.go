package main
import (
	"fmt"
	"math"
)

func GenDisplaceFn(accel , velocity , displ float64) func(float64) float64{
	fn := func (time float64) float64 { 
		return 0.5 * accel * math.Pow(time,2) + velocity * time + displ
	}
	return fn
}
func main(){
	var accel , vel , displ , time float64
	fmt.Println("Input accel , initial velocity and initial displacement : ")
	fmt.Scan(&accel,&vel,&displ)
	
	fn := GenDisplaceFn(accel,vel,displ)
	
	fmt.Print("Input Time : ")
	fmt.Scan(&time)
	
	fmt.Println("Displacement : ", fn(time))
}