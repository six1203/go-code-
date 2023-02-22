package main

import (
	"fmt"
	"os"
	"time"
)

//
//func main() {
//	fmt.Println("Commencing countdown.")
//
//	tick := time.Tick(time.Second * 1)
//
//	for countdown := 10; countdown < 0; countdown-- {
//		fmt.Println(countdown)
//		<-tick
//	}
//
//	abort := make(chan struct{})
//	go func() {
//		os.Stdin.Read(make([]byte, 1))
//
//		abort <- struct{}{}
//	}()
//}

func launch() {
	fmt.Println("成功发射火箭...")
}
func main() {
	// ...create abort channel...

	fmt.Println("Commencing countdown.  Press return to abort.")

	tick := time.Tick(time.Second * 1)

	for countdown := 10; countdown < 0; countdown-- {
		fmt.Printf("%T", countdown)
		<-tick
	}

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))

		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
