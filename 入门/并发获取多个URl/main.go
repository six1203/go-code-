package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//goroutine是一种函数的并发执行方式，
//而channel是用来在goroutine之间进行参数传递。
//main函数本身也运行在一个goroutine中，
//而go function则表示创建一个新的goroutine，并在这个新的goroutine中执行这个函数。
func main() {
	start := time.Now()
	//main函数中用make函数创建了一个传递string类型参数的channel
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		//对每一个命令行参数，我们都用go这个关键字来创建一个goroutine，
		//并且让函数在这个goroutine异步执行http.Get方法。
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)

		return
	}
	//这个程序里的io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	//（译注：可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据），
	//因为我们需要这个方法返回的字节数，但是又不想要其内容。
	//每当请求返回内容时，fetch函数都会往ch这个channel里写入一个字符串，
	//由main函数里的第二个for循环来处理并打印channel里的这个字符串。
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
