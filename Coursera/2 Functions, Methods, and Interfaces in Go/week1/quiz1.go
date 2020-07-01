package main

import "fmt"

func Swap(arr []int, index int){
	arr[index] , arr[index+1] = arr[index+1] , arr[index]
}

func BubbleSort(arr []int) {
	
	// Tracks changes to the slice
	flag_changed := true
	for flag_changed == true {
		
		// No changes initially 
		flag_changed = false

		for i, _ := range arr {

			// Skip firs element
			if i == 0 {
				continue
			}

			// Compare current element with previous element
			if arr[i] < arr[i-1] {

				Swap(arr,i-1) // Swap with previous
				flag_changed = true // Change occurred in slice
				
			}
		}
	}
}

func main() {
	
	fmt.Println("Input a list of 10 integers")
	
	// Input list elements
	var list []int
	for i:=0 ; i<10 ; i++ {
		var tmp int
		fmt.Scan(&tmp)
		list = append(list,tmp)
	}

	// Sort
	BubbleSort(list)

	//Output
	fmt.Println("\nSorted list : " , list)
}
