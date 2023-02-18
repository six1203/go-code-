package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	//在这些代码的背后，服务器每一次接收请求处理时都会另起一个goroutine，这样服务器就可以同一时间处理多个请求。
	//然而在并发情况下，假如真的有两个请求同一时刻去更新count，那么这个值可能并不会被正确地增加；
	//我们必须保证每次修改变量的最多只能有一个goroutine，
	//这也就是代码里的mu.Lock()和mu.Unlock()调用将修改count的所有行为包在中间的目的。
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(
		http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 线程安全
	//mu.Lock()
	//count++
	//mu.Unlock()
	// Fprintf 标注输出流

	fmt.Fprintf(w, "%s %s %s %s\n", r.Method, r.URL, r.Header, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "Remoteadd = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "Referer = %q\n", r.Referer())
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]=%q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
