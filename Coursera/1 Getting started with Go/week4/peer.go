package main
import ("encoding/json"
"fmt")
func main(){
var name,addr string
fmt.Scan(&name)
fmt.Scan(&addr)
person:=map[string] string{"name":name,"address":addr}
object,_:=json.Marshal(person)
fmt.Println(string(object))
}