package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test() hello" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

var (
	resMap = make(map[int]int)
	lock   sync.Mutex
)

func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	resMap[n] = res
	lock.Unlock()
}

func main() {
	//go test()
	//
	//for i := 1; i <= 10; i++ {
	//	fmt.Println("main() hello" + strconv.Itoa(i))
	//	time.Sleep(time.Second)
	//}
	//num := runtime.NumCPU()
	//runtime.GOMAXPROCS(num)
	//fmt.Println(num)

	for i := 1; i < 200; i++ {
		go factorial(i)
	}
	time.Sleep(time.Second * 10)

	lock.Lock()
	for key, v := range resMap {
		fmt.Println(key, v)
	}
	lock.Unlock()

}
