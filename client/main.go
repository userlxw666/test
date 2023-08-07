package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("conn err", err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			_, err = conn.Write([]byte("hello server"))
			if err != nil {
				fmt.Println("read err", err)
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

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
