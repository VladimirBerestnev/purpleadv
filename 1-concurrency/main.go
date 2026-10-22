package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	slice := []int{}
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go Write(ch)
	go func() {
		num := <-ch
		sqrt := int(math.Pow(float64(num), 2))
		slice = append(slice, sqrt)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(slice)
}

func Write(ch chan int) {
	slice := []int{}
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		slice = append(slice, num)
	}
	fmt.Println("Write:", slice)
	for _, item := range slice {
		ch <- item
	}
}
