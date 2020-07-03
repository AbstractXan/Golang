package main
import "fmt"
type Animal struct{
	food string
	locomotion string
	noise string
}
func (a Animal) Eat(){
	fmt.Println(a.food)
}
func (a Animal) Move(){
	fmt.Println(a.locomotion)
}
func (a Animal) Speak(){
	fmt.Println(a.noise)
}

func getRequest(animalList []Animal){
	fmt.Print("> ")
	
	var animal, action string
	fmt.Scanf("%s %s", &animal ,&action)

	var index int8
	switch animal{
	case "cow":
		index = 0
	case "bird":
		index = 1
	case "snake":
		index = 2
	default: 
		fmt.Println("Oops! We only have a cow, a bird and a snake")
		return
	}

	switch action{
	case "eat":
		animalList[index].Eat()
	case "move":
		animalList[index].Move()
	case "speak":
		animalList[index].Speak()
	default: 
		fmt.Println("Oops! our animals could only 'eat', 'move' or 'speak")
		return
	}
}
func main(){
	cow := Animal{food:"grass", locomotion:"walk", noise:"moo"}
	bird := Animal{food:"worms", locomotion:"fly", noise:"peep"}
	snake := Animal{food:"mice", locomotion:"slither", noise:"hsss"}
	animalList := []Animal{cow, bird, snake}

	for {
		getRequest(animalList)
	}
	
}