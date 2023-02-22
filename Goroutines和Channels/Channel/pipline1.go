package main

import (
	"fmt"
)

func counter(out chan int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	//naturals := make(chan int)
	//squares := make(chan int)

	//go func() {
	//	for x := 0; x < 100; x++ {
	//		naturals <- x
	//	}
	//	close(naturals)
	//}()
	//
	//go func() {
	//	for x := range naturals {
	//		squares <- x * x
	//	}
	//	close(squares)
	//}()
	//
	//for x := range squares {
	//	fmt.Println(x)
	//}
	//go counter(naturals)
	//go squarer(naturals, squares)
	//printer(squares)

	x := make(chan int, 3)

	x <- 1
	x <- 2
	x <- 3
	fmt.Println(<-x)
	fmt.Println(<-x)
	fmt.Println(<-x)
	fmt.Println(<-x)
}

//func request(hostname string) (response string) { /* ... */ }
