package main
import(
"fmt"
"time"
)
func main(){
	go count("sheep")
	go count("fish")

	// If main function finishes, it exits regardless of goroutines
	
	// Methods to prevent:
	// 1. Sleep
	// time.Sleep(time.Second * 2)
	// 2. Scanline
	//fmt.Scanln()
	// 3. Wait group
	// Refer to next code example file
}

func count(thing string){
	for i := 1; true; i++{
		fmt.Println(i,thing)
		time.Sleep(time.Millisecond*500)
	}
}
