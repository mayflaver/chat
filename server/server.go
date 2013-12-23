package main

import (
	"net"
	"fmt"
)

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		// handle error
	}
	c_conn, d_conn, c_msg := make(chan net.Conn), make(chan net.Conn), make(chan []byte)
	conns := make(map[net.Conn]int)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				// handle error
				continue
			}
			c_conn <- conn
		}
	}()
	for {
		select {
		case conn := <- c_conn:
			conns[conn] = 0
			fmt.Println("connect: ", conn.RemoteAddr())
			go func(conn net.Conn) {
				conn.Write([]byte("welcome"))
				for {
					buf := make([]byte, 64)
					_, err := conn.Read(buf)
					if err != nil {
						d_conn <- conn
						return
					}
					c_msg <- buf
				}
			}(conn)
		case msg := <- c_msg:
			for conn, _ := range conns {
				conn.Write(msg)
			}
		case conn := <- d_conn:
			fmt.Println("disconnect: ", conn.RemoteAddr())
			delete(conns, conn)
			
			
		}
	}
}

















