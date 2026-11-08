package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	ch := make(chan int)
	ch2 := make(chan int)
	//var wg sync.WaitGroup

	go Write(ch)
	go func() {
		for num := range ch {
			sqrt := int(math.Pow(float64(num), 2))
			ch2 <- sqrt
		}
		close(ch2)
	}()

	for res := range ch2 {
		fmt.Println(res)
	}

}

func Write(ch chan int) {
	defer close(ch)
	slice := []int{}
	for i := 0; i < 10; i++ {
		num := rand.Intn(101)
		slice = append(slice, num)
	}
	fmt.Println("Write:", slice)
	for _, item := range slice {
		ch <- item
	}
}
