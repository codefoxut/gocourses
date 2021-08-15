package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MAX_TIMES_ALLOWED = 3

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	num, counter      int
	leftCS, rightCS   Chopstick
	permission_to_eat bool
}

func (p *Philosopher) Eat() {
	if p.counter < MAX_TIMES_ALLOWED {
		choose_cs := rand.Intn(2)
		if choose_cs == 0 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}

		fmt.Printf("starting to eat <%d>\n", p.num)
		p.counter++
		fmt.Printf("finishing eating <%d>. eaten for %d time \n", p.num, p.counter)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
		p.reset_permission()

	} else {
		fmt.Println(p.num, "is already done.")
	}

}

func (p *Philosopher) allow_to_eat() {
	p.permission_to_eat = true
}

func (p *Philosopher) reset_permission() {
	p.permission_to_eat = false
}

func (p *Philosopher) AllowedToEat() bool {
	return p.permission_to_eat
}

func host(ch chan int, ps []*Philosopher) {
	defer close(ch)
	rand.Seed(time.Now().UnixNano())
	var x, y int
	x = rand.Intn(len(ps))
	for ps[x].AllowedToEat() || ps[x].counter >= MAX_TIMES_ALLOWED {
		x = rand.Intn(len(ps))
	}
	ch <- x

	y = rand.Intn(len(ps))
	for (x == y || ps[y].AllowedToEat() || ps[y].counter >= MAX_TIMES_ALLOWED) && is_dinner_going_on(ps) {
		y = rand.Intn(len(ps))
	}

	if is_dinner_going_on(ps) {
		ch <- y
	}
	fmt.Println("philosophers choosen are ", x+1, y+1)
}

func is_dinner_going_on(ps []*Philosopher) bool {
	for _, p := range ps {
		// fmt.Println(p)
		if p.counter < MAX_TIMES_ALLOWED {
			return true
		}
	}
	return false
}
func can_ask_for_host_permission(ps []*Philosopher) bool {
	eating_philos := 0
	for _, p := range ps {
		if p.AllowedToEat() {
			eating_philos++
		}
	}
	// fmt.Println(eating_philos)
	if eating_philos > 1 {
		return false
	}
	return true
}

func main() {
	start := time.Now()
	CSticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chopstick)
		fmt.Println(*CSticks[i])
	}

	Philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{i + 1, 0, *CSticks[i], *CSticks[(i+1)%5], false}
		fmt.Println(*Philosophers[i])
	}

	for is_dinner_going_on(Philosophers) {
		if can_ask_for_host_permission(Philosophers) {
			host_channel := make(chan int, 2)
			go host(host_channel, Philosophers)
			for philo := range host_channel {
				Philosophers[philo].allow_to_eat()
				go Philosophers[philo].Eat()
			}

		}

	}

	fmt.Println(time.Since(start))

}
