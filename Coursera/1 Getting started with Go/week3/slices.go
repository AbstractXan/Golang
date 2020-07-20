package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sli := make([]int,0,3)

	// Input rest of the numbers
	for {
		fmt.Print("Enter a number to be added to slice \n> ")
		var tmp int
		fmt.Scan(&tmp)
		if tmp==0 {
			return
		}
		sli = append(sli, tmp)
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
