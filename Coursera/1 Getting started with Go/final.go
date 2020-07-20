package main
import (
	"fmt"
	"os"
	"log"
	"bufio"
)

type name struct{
	fname string
	lname string
}

func scanfile(file *os.File) []name {
	scanner := bufio.NewScanner(file)
	var names []name
	for scanner.Scan() {
		var firstName , secondName string
		_, err := fmt.Sscanf(scanner.Text(), "%s %s", &firstName, &secondName)
		if err != nil {
			log.Fatal(err)
		}

		names = append(names,name{fname: firstName, lname: secondName})
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}

	return names
}
func main(){
	fmt.Println("Enter the name of the file (example: file.txt)")
	var filename string

	fmt.Scan(&filename)
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	names := scanfile(file)
	for i,elem := range names{
		fmt.Println(i,"Fistname:",elem.fname, "Lastname:",elem.lname)
	}
}
