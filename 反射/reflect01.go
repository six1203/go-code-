package main

import (
	"fmt"
	"reflect"
)

func test(a interface{}) {
	rType := reflect.TypeOf(a)

	fmt.Printf("rType=%v \n", rType)

	rVal := reflect.ValueOf(a)
	fmt.Printf("rVal=%v rType=%T \n", rVal, rVal)

	//rVal.Int()

	iv := rVal.Interface()

	if num2, ok := iv.(int); ok {
		fmt.Println(num2)
	} else {
		fmt.Println("convert failed")
	}

}

func main() {
	var num int = 100
	test(num)
}
