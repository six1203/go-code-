package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

//func ReadFile(filename string) ([]byte, error) {
//	f, err := os.Open(filename)
//	if err != nil {
//		return nil, err
//	}
//	defer f.Close()
//
//	return ReadALL(f)
//}
var mu sync.Mutex

var m = make(map[string]int)

func loop(key string) int {
	mu.Lock()
	defer mu.Unlock()

	return m[key]
}

func bigSlowOperation() {
	// defer机制也常被用于记录何时进入和退出函数
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	//bigSlowOperation()
	double(5)
}

func double(x int) (result int) {
	defer func() { fmt.Println(x, result) }()
	return x * x
}
