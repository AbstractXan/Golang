package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	animals := make(map[string]Animal)

	fmt.Println("Create an animal with 'newanimal <name> <cow/bird/snake>' or query a created animal with 'query <name> <eat/move/speak>'. Exit with 'q'.")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		input, inputErr := reader.ReadString('\n')
		if inputErr != nil {
			fmt.Printf("Invalid Input. Exiting...\n")
			os.Exit(1)
		}

		input = strings.TrimRight(strings.Replace(input, "\r\n", "\n", -1), "\n")
		if input == "q" {
			break
		}

		tokens := strings.Split(input, " ")
		if len(tokens) != 3 {
			fmt.Println("Expected three arguments.")
			continue
		}

		switch tokens[0] {
		case "newanimal":
			name := tokens[1]
			animalType := tokens[2]
			_, exists := animals[name]
			if exists {
				fmt.Printf("An animal with the name %s already exists.\n", name)
				continue
			}
			switch animalType {
			case "cow":
				animals[name] = Cow{name}
			case "bird":
				animals[name] = Bird{name}
			case "snake":
				animals[name] = Snake{name}
			default:
				fmt.Println("Third argument has to be cow/bird/snake.")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			name := tokens[1]
			action := tokens[2]
			_, exists := animals[name]
			if !exists {
				fmt.Printf("No animal with the name %s has been created.\n", name)
				continue
			}
			switch action {
			case "eat":
				animals[name].Eat()
			case "move":
				animals[name].Move()
			case "speak":
				animals[name].Speak()
			default:
				fmt.Println("Third argument has to be eat/move/speak.")
				continue
			}
		default:
			fmt.Println("First argument must either be 'newanimal' or 'query'.")
		}
	}

	fmt.Println("Bye!")
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Printf("%s eats grass\n", c.name)
}

func (c Cow) Move() {
	fmt.Printf("%s walks\n", c.name)
}

func (c Cow) Speak() {
	fmt.Printf("%s moos\n", c.name)
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Printf("%s eats worms\n", b.name)
}

func (b Bird) Move() {
	fmt.Printf("%s flies\n", b.name)
}

func (b Bird) Speak() {
	fmt.Printf("%s peeps\n", b.name)
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Printf("%s eats mice\n", s.name)
}

func (s Snake) Move() {
	fmt.Printf("%s slithers\n", s.name)
}

func (s Snake) Speak() {
	fmt.Printf("%s hisses\n", s.name)
}