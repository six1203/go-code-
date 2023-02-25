package main

import "fmt"

type usb interface {
	start()
	stop()
}

type camera struct {
}

type iphone struct {
}

type computer struct {
}

func (c camera) start() {
	fmt.Println("相机开始工作了")
}

func (c camera) stop() {
	fmt.Println("相机暂停工作了")
}

func (p iphone) start() {
	fmt.Println("手机开始工作了")
}

func (p iphone) stop() {
	fmt.Println("手机开始工作了")
}

func (c computer) working(u usb) {
	u.start()
	u.stop()
}

func main() {
	c := computer{}
	p := iphone{}
	ca := camera{}

	c.working(p)
	c.working(ca)
}
