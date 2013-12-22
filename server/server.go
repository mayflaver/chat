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
	c_conn, c_msg := make(chan net.Conn), make(chan []byte)
	conns := make([]net.Conn, 0)
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
			conns = append(conns, conn)
			fmt.Println(conns)
			go func(conn net.Conn) {
				conn.Write([]byte("welcome"))
				for {
					buf := make([]byte, 64)
					conn.Read(buf)
					c_msg <- buf
				}
			}(conn)
		case msg := <- c_msg:
			for _, conn := range conns {
				fmt.Println(msg)
				fmt.Println(conn)
				conn.Write(msg)
			}
		}
	}
}
