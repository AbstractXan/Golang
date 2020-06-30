package main
import "fmt"
func main(){
	for i := 0; i< 5;i++{
		fmt.Println(i)
	}

	arr := []string{"a","b","c"}
	
	for index, value := range arr {
		fmt.Println("index" , index, "value", value)
	}
}
