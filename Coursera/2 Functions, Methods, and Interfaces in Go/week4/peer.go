package main

import (
	"fmt"
	"log"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name, food, locomotion, noise string
}

func (c *Cow) Eat() {
	fmt.Println(c.food)
}

func (c *Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c *Cow) Speak() {
	fmt.Println(c.noise)
}

func NewCow(name string) *Cow {
	return &Cow{
		name:       name,
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
}

type Bird struct {
	name, food, locomotion, noise string
}

func (b *Bird) Eat() {
	fmt.Println(b.food)
}

func (b *Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b *Bird) Speak() {
	fmt.Println(b.noise)
}

func NewBird(name string) *Bird {
	return &Bird{
		name:       name,
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
}

type Snake struct {
	name, food, locomotion, noise string
}

func (s *Snake) Eat() {
	fmt.Println(s.food)
}

func (s *Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s *Snake) Speak() {
	fmt.Println(s.noise)
}

func NewSnake(name string) *Snake {
	return &Snake{
		name:       name,
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
}

func NewAnimal(animalName string, animalType string) (Animal, error) {
	animalType = strings.ToLower(animalType)
	switch animalType {
	case "cow":
		return NewCow(animalName), nil
	case "bird":
		return NewBird(animalName), nil
	case "snake":
		return NewSnake(animalName), nil
	default:
		return nil, fmt.Errorf("unknown animal type %s", animalType)
	}
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type UserInput struct {
	command, arg1, arg2 string
}

func getUserInput() *UserInput {
	fmt.Println("================================")
	fmt.Println("Enter one of command:")
	fmt.Println("newanimal name [cow|bird|snake]. Example: newanimal Molly cow ")
	fmt.Println("query name [eat|move|speak]. Example: query Molly move")
	fmt.Print("> ")

	var command, arg1, arg2 string
	_, err := fmt.Scanf("%s %s %s\n", &command, &arg1, &arg2)
	fatalOnError(err)
	command = strings.ToLower(command)
	return &UserInput{
		command: command,
		arg1:    arg1,
		arg2:    arg2,
	}
}

func ProcessInput(input *UserInput, db map[string]Animal) {
	switch input.command {
	case "newanimal":
		animalName, animalType := input.arg1, input.arg2

		if animal, err := NewAnimal(animalName, animalType); err == nil {
			animalName = strings.ToLower(animalName)
			db[animalName] = animal
			fmt.Println("Created it!")
		} else {
			fmt.Println(err)
		}

	case "query":
		animalName, queryType := input.arg1, input.arg2
		animalName = strings.ToLower(animalName)
		queryType = strings.ToLower(queryType)

		if animal, ok := db[animalName]; ok {
			switch queryType {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Printf("unknown query type %s\n", queryType)
			}
		} else {
			fmt.Printf("unknown name %s\n", animalName)
		}
	default:
		fmt.Printf("unknown command %s\n", input.command)
	}
	fmt.Println()
}

func main() {
	db := map[string]Animal{}
	for {
		input := getUserInput()
		ProcessInput(input, db)
		fmt.Println(db)
	}
}