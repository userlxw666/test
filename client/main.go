package main

import (
	"fmt"
	"net"
	"test1/pack"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("conn err", err)
		return
	}
	defer conn.Close()
	dp := pack.NewPack()
	msg1 := pack.NewMessage(1, []byte("send task1 server"))
	msg2 := pack.NewMessage(2, []byte("send task2 server"))
	msg1Pack, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("pack msg1 error", err)
		return
	}
	msg2Pack, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("pack msg2 error", err)
		return
	}
	sendMsg := append(msg1Pack, msg2Pack...)
	// send server msg
	go func() {
		for {
			_, err = conn.Write(sendMsg)
			if err != nil {
				fmt.Println("read err", err)
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()
	// read server message
	go func() {
		for {
			var buf = make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read err", err)
				return
			}
			fmt.Println("rec server back msg :", string(buf[:n]))

		}
	}()
	select {}
}
