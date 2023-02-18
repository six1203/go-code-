package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	var step string
	for i := 1; i < len(os.Args); i++ {

		s += step + os.Args[i]
		step = " "

	}

	fmt.Println(strings.Join(os.Args[1:], ""))

	str := "赵,钱,孙,李,赵"
	str1 := strings.Split(str, ",")
	fmt.Printf(str1[0])

	fmt.Println(s)
}
