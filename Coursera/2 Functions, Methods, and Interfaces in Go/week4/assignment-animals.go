package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}
type animal struct {
	name string
}

type Cow animal
type Bird animal
type Snake animal

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

var animalList []Animal

func runQuery(name, action string) {
	var anim Animal
	for _, curr_animal := range animalList {
		curr1, ok := curr_animal.(Cow)
		if ok {
			if curr1.name == name {
				anim = curr_animal
				break
			}
			continue
		}

		curr2, ok := curr_animal.(Snake)
		if ok {
			if curr2.name == name {
				anim = curr_animal
				break
			}
			continue
		}

		curr3, ok := curr_animal.(Bird)
		if ok {
			if curr3.name == name {
				anim = curr_animal
				break
			}
			continue
		}
	}

	if anim != nil {
		switch action {
		case "eat":
			anim.Eat()
		case "move":
			anim.Move()
		case "speak":
			anim.Speak()
		}
	}
}
func runNewAnimal(name, animalType string) {
	var anim Animal
	switch animalType {
	case "cow":
		anim = Cow{name: name}
	case "bird":
		anim = Bird{name: name}
	case "snake":
		anim = Snake{name: name}
	}

	animalList = append(animalList, anim)
	fmt.Println("Created it!")
}

func getRequest() {
	fmt.Print("> ")

	var command, str1, str2 string
	fmt.Scanf("%s %s %s", &command, &str1, &str2)

	switch command {
	case "newanimal":
		runNewAnimal(str1, str2)
	case "query":
		runQuery(str1, str2)
	}

}

func main() {
	fmt.Println("Input format : \nnewanimal <name> cow/snake/bird \nquery <name> eat/move/speak")
	for {
		getRequest()
	}
}
