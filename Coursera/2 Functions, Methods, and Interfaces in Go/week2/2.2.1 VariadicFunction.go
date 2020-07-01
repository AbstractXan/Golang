package main
import "fmt"

func getMax(vals ...int) int {
	maxVal := -1

	// Treated like a slice
	for _,v := range vals{
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
func main(){
	vslice := []int{1,2,3,4,10}
	fmt.Println(getMax(vslice...))
}