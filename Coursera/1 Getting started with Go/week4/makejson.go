package main
import (
	"fmt"
	"encoding/json"
)

func main(){
	var name , address string
	fmt.Scan(&name, &address)

	m := make(map[string] string)
	m["name"] = name
	m["address"] = address
	data , _ := json.Marshal(m)
	fmt.Println(string(data))
}