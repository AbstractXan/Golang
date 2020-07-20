package main

import (
	"fmt"
	"sync"
	"time"
)

type chop struct{ sync.Mutex }

type philo struct {
	id                  int
	leftchop, rightchop *chop
}

func (p philo) eat() {
	for i := 0; i < 3; i++ {
		p.leftchop.Lock()
		p.rightchop.Lock()

		print("eating", p.id)
		time.Sleep(time.Second)

		p.rightchop.Unlock()
		p.leftchop.Unlock()

		print(" eating finished", p.id)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func print(action string, id int) {
	fmt.Printf("philo %d is %s\n", id+1, action)
}

var wg sync.WaitGroup

func main() {

	chops := make([]*chop, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(chop)
	}

	philos := make([]*philo, 5)
	fmt.Print(philos)
	for i := 0; i < 5; i++ {
		philos[i] = &philo{
			id: i, leftchop: chops[i], rightchop: chops[(i+1)%5]}
		wg.Add(1)
		go philos[i].eat()

	}
	wg.Wait()
}