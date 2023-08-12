package main

import (
	"fmt"
	"net"
	"test1/pack"
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
		// reciver client msg
		go StartRead(conn)
		// send client msg
		go SendMsg(conn)
	}
}

func SendMsg(conn net.Conn) {
	for {
		select {
		case <-writeChan:
			data := fmt.Sprintf("server receiver message，content ：%s", "completed task")
			_, err := conn.Write([]byte(data))
			if err != nil {
				fmt.Println("send client msg error", err)
				return
			}
		}
	}
}

func StartRead(conn net.Conn) {
	for {
		dp := pack.NewPack()
		dataHead := make([]byte, dp.GetHeadLen())
		_, err := conn.Read(dataHead)
		if err != nil {
			fmt.Println("read head msg error", err)
			return
		}
		msghead, err := dp.UnPack(dataHead)
		if err != nil {
			fmt.Println("unpack headmsg error", err)
			return
		}

		msg := msghead.(*pack.Message)
		msg.Data = make([]byte, msg.DataLen)
		_, err = conn.Read(msg.Data)
		if err != nil {
			fmt.Println("read data error", err)
			return
		}
		req := &Request{
			conn: conn,
			msg:  msg.Data,
		}
		SendToQueue(req)
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
