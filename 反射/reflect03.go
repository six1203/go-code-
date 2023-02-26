package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// addr 小写字段是非导出字段，不可以从包外部访问
	addr string `json:"_"`
}

func main() {
	u := User{1, "dylan", "上海市"}

	rT := reflect.TypeOf(u)

	rV := reflect.ValueOf(u)

	kind := rV.Kind()

	numfield := rV.NumField()

	fmt.Printf("rt=%v\n", rT)
	fmt.Printf("rv=%v\n", rV)
	fmt.Printf("rt=%v\n", kind)
	fmt.Printf("rt=%v\n", numfield)
	fmt.Printf("rs=%v\n", reflect.Struct)

	for i := 0; i < rV.NumField(); i++ {
		if rV.Field(i).CanInterface() {
			fmt.Printf("%s %s = %v - tag:%s\n", rT.Field(i).Name, rT.Field(i).Type, rV.Field(i).Interface(), rT.Field(i).Tag)
		}
	}
}
