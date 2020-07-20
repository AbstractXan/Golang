package main

import (
	"fmt"
	"sync"
	"time"
)

type stick struct{ sync.Mutex }

type philosopher struct {
	id                  int
	leftStick, rightStick *stick
}

func (p philosopher) eat() {
	for j := 0; j < 3; j++ {
		p.leftStick.Lock()
		p.rightStick.Lock()

		fmt.Printf("Starting to eat %d\n", p.id+1)
		time.Sleep(time.Second)


		fmt.Printf("Finishing eating %d\n", p.id+1)
		time.Sleep(time.Second)

		p.rightStick.Unlock()
		p.leftStick.Unlock()

	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	count := 5

	chopsticks := make([]*stick, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = new(stick)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, leftStick: chopsticks[i], rightStick: chopsticks[(i+1)%count]}
		wg.Add(1)
		go philosophers[i].eat()

	}
	wg.Wait()

}