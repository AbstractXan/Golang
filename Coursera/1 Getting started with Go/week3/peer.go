package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	sl := make([]int, 0, 3)
	for {
		fmt.Printf("Enter one integer (Enter X to exit): ")
		input.Scan()
		s := input.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			if s == "X" {
				break
			}
		}
		sl = append(sl, i)
		sort.Ints(sl)
		for _, n := range sl {
			fmt.Printf("%d ", n)
		}
		fmt.Printf("\n")
	}
}