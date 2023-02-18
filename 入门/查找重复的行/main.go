// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//func main() {
//	counts := make(map[string]int)
//	input := bufio.NewScanner(os.Stdin)
//
//	//该变量从程序的标准输入中读取内容。每次调用 input.Scan()，即读入下一行，并移除行末的换行符；
//	//读取的内容可以调用 input.Text() 得到。
//	//Scan 函数在读到一行时返回 true，不再有输入时返回 false。
//	for input.Scan() {
//		//line := input.Text()
//		//counts[line] = counts[line] + 1
//		counts[input.Text()]++
//	}
//	// NOTE: ignoring potential errors from input.Err()
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}

//func main() {
//	counts := make(map[string]int)
//	files := os.Args[1:]
//	if len(files) == 0 {
//		countLines(os.Stdin, counts)
//	} else {
//		for _, arg := range files {
//			// os.Open 函数返回两个值。第一个值是被打开的文件（*os.File），其后被 Scanner 读取。
//			//os.Open 返回的第二个值是内置 error 类型的值。如果 err 等于内置值nil
//			//（译注：相当于其它语言里的 NULL），那么文件被成功打开。读取文件，直到文件结束，
//			//然后调用 Close 关闭该文件，并释放占用的所有资源。
//			//相反的话，如果 err 的值不是 nil，说明打开文件时出错了。这种情况下，错误值描述了所遇到的问题。我们的错误处理非常简单，
//			//只是使用 Fprintf 与表示任意类型默认格式值的动词 %v，向标准错误流打印一条信息，
//			//然后 dup 继续处理下一个文件；continue 语句直接跳到 for 循环的下个迭代开始执行。
//			f, err := os.Open(arg)
//			if err != nil {
//				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
//				continue
//			}
//
//			countLines(f, counts)
//		}
//	}
//
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}
//
//func countLines(f *os.File, counts map[string]int) {
//	//map 是一个由 make 函数创建的数据结构的引用。map 作为参数传递给某函数时，该函数接收这个引用的一份拷贝（copy，或译为副本），
//	//被调用函数对 map 底层数据结构的任何修改，调用者函数都可以通过持有的 map 引用看到。
//	//在我们的例子中，countLines 函数向 counts 插入的值，也会被 main 函数看到。
//	//（译注：类似于 C++ 里的引用传递，实际上指针是另一个指针了，但内部存的值指向同一块内存）
//	input := bufio.NewScanner(f)
//
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//}

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
