package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	in := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string")
	in.Scan()

	s := strings.ToLower(in.Text())

	foundA := strings.Contains(s,"a")

	if s[0] == 'i' && s[len(s)-1] == 'n' && foundA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
