package main

import (
	"fmt"
)

func Swap(arr []int, index int) {
	arr[index], arr[index+1] = arr[index+1], arr[index]
}

func BubbleSort(goroutineNum int , arr []int, c chan []int) {
	fmt.Println(goroutineNum, "Sorting: ", arr)
	flag_changed := true
	for flag_changed == true {
		flag_changed = false
		for i, _ := range arr {
			// Skip first element
			if i == 0 {
				continue
			}
			// Compare current element with previous element
			if arr[i] < arr[i-1] {
				Swap(arr, i-1)      // Swap with previous
				flag_changed = true // Change occurred in slice
			}
		}
	}
	c <- arr
}

func Merge(sorted, sli []int) []int {
	sorted_index := 0
	sli_index := 0
	sortedLength, sliLength := len(sorted), len(sli)

	var merged []int
	for sorted_index != sortedLength && sli_index != sliLength {
		if sorted[sorted_index] <= sli[sli_index] {
			merged = append(merged, sorted[sorted_index])
			sorted_index += 1
		} else {
			merged = append(merged, sli[sli_index])
			sli_index += 1
		}
	}

	if sorted_index == sortedLength {
		merged = append(merged, sli[sli_index:sliLength]...)

	} else {
		merged = append(merged, sorted[sorted_index:sortedLength]...)
	}
	
	return merged
}
func main() {
	fmt.Println("Enter the number of numbers (atleast 4)")
	var num int
	fmt.Scan(&num)

	if num < 4 {
		fmt.Println(" Number of numbers can't be less than 4 ")
		return
	}
	// Make slice
	var list []int
	for i := 0; i < num; i++ {
		var temp int
		fmt.Scan(&temp)
		list = append(list, temp)
	}

	// Split the slice
	diff := len(list) / 4
	index := 0

	c := make(chan []int, 4)
	for i := 0; i < 3; i++ {
		go BubbleSort(i,list[index:index+diff], c)
		index += diff
	}

	// Last slice may need to have extra elemets
	go BubbleSort(3,list[index:len(list)],c)

	var sorted []int
	for i := 0; i < 4; i++ {
		sli := <-c
		sorted = Merge(sorted, sli)
	}

	fmt.Println(sorted)
}
