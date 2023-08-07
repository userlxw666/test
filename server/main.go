package main

import (
	"fmt"
	"net"
)

type Request struct {
	conn net.Conn
	msg  []byte
}

var clients = make(map[net.Conn]int)

var queueChan = make([]chan []byte, 3)

var writeChan = make(chan []byte)

var a = 0

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("listener error", err)
		return
	}
	defer listener.Close()
	go WorkPool()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connect error", err)
			return
		}
		clients[conn] = a
		a++
		go func() {
			for {
				var buf = make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}
				req := &Request{
					conn: conn,
					msg:  buf[:n],
				}
				SendToQueue(req)
			}

		}()
		go func() {
			for {
				select {
				case msg := <-writeChan:
					data := fmt.Sprintf("server receiver message，content ：%s", string(msg))
					_, err := conn.Write([]byte(data))
					if err != nil {
						fmt.Println("send client msg error", err)
						return
					}
				}
			}
		}()
	}
}

func WorkPool() {
	for i := 0; i < 3; i++ {
		queueChan[i] = make(chan []byte, 1024)
		go StartWork(i, queueChan[i])
	}
}

func SendToQueue(req *Request) {
	num := clients[req.conn] % 3
	queueChan[num] <- req.msg
	writeChan <- req.msg

}

func StartWork(id int, queueChan chan []byte) {
	fmt.Println("worker=", id, "is starting")
	for {
		select {
		case msg := <-queueChan:
			fmt.Println("worker", id, "rec msg", string(msg))
		}
	}

}
