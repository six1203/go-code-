package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)

	leaving = make(chan client)

	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

/* broadcaster监听来自全局的entering和leaving的channel来获知客户端的到来和离开事件
当其接收到其中的一个事件时，会更新clients集合，当该事件是离开行为时，它会关闭客户端的消息发送channel。
broadcaster也会监听全局的消息channel，所有的客户端都会向这个channel中发送消息。
当broadcaster接收到什么消息时，就会将其广播至所有连接到服务端的客户端。*/
func broadcaster() {
	clients := make(map[client]bool)
	for {

		/*select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
		select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
		如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。*/
		select {

		// 从messages channel中取消息
		case msg := <-messages:
			// 向所有人广播消息
			for cli := range clients {
				cli <- msg
			}

			// 有客户端连接上就加入到clients map中
		case cli := <-entering:

			clients[cli] = true

			// 有客户端离开，从客户端map中删除，并且关闭channel
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

/* handleConn函数会为它的客户端创建一个消息发送channel并通过entering channel来通知客户端的到来。
然后它会读取客户端发来的每一行文本，并通过全局的消息channel来将这些文本发送出去，
并为每条消息带上发送者的前缀来标明消息身份。当客户端发送完毕后，handleConn会通过leaving这个channel来通知客户端的离开并关闭连接。
*/
func handleConn(conn net.Conn) {
	ch := make(chan string)
	/*另外，handleConn为每一个客户端创建了一个clientWriter的goroutine，
	用来接收向客户端发送消息的channel中的广播消息，并将它们写入到客户端的网络连接。
	客户端的读取循环会在broadcaster接收到leaving通知并关闭了channel后终止。*/
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	ch <- "you are " + who

	messages <- who + "has arrived"

	entering <- ch

	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- who + ":" + input.Text()
	}

	leaving <- ch

	messages <- who + "has left"

	conn.Close()
}

func clientWriter(conn net.Conn, ch chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
