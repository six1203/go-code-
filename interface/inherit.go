package main

import "fmt"

type AInterface interface {
	hello()
	test()
}

type BInterface interface {
	say()
	test()
}

type CInterface interface {
	AInterface
	BInterface
}

type student struct {
}

func (stu student) say() {

}

func (stu student) test() {

}
func (stu student) hello() {

}

func main() {
	stu := student{}
	var a CInterface = stu
	a.say()
	a.test()
	a.hello()

	b := "我是好的好的和"
	c := "哈哈"

	fmt.Println(b < c)
}
