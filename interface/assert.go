package main

import "fmt"

func typeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("... %d, %T", i, x)
		case float64:
			fmt.Printf("... %d, %T ", i, x)

		case int, int64:
			fmt.Printf("... %d, %T", i, x)

		default:
			fmt.Printf("... %d, %T", i, x)
		}
	}
}

func main() {
	var n1 float64 = 1.11
	var n2 bool = false
	var n3 int16 = 3
	var n4 string = "string"

	typeJudge(n1, n2, n3, n4)
}
