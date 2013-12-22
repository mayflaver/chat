package main

import (
	"fmt"
	"net"
	"bytes"
)

func main() {
	from_server, to_server, console, quit:= make(chan []byte), make(chan []byte), make(chan []byte), make(chan string)
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		// handle error
	}
	go func(conn net.Conn) {
		for {
			buf := make([]byte, 64)
			_, err := conn.Read(buf)
			if err != nil {
				quit <-"quit"
				return
			}
			from_server <- buf
		}
	}(conn)
	go func(conn net.Conn) {
		for {
			buf := <- to_server
			conn.Write(buf)
		}
	}(conn)
	go func() {
		for  {
			str := ""
			_, err := fmt.Scanln(&str)
			if err != nil {
				// handle error
			}
			console <- []byte(str)
		}
	}()
	for {
		select {
		case str := <-from_server:
			n := bytes.Index(str, []byte{0})
			s := string(str[:n])
			fmt.Println(s)
		case str := <-console:
			to_server <- str
		case <-quit:
			return
		}
	}
}
	
		


		












