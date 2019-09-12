package main

import (
	//"bufio"
	"fmt"
	"io"
	"net"
	//"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		fmt.Print("err:", err)
	} else {
		fmt.Printf("%#v\n", conn)
		for {
			request := "GET / HTTP/1.1\r\nHost:ipaddress.com\r\n\r\n"
			conn.Write([]byte(request))
			all := make([]byte, 0)   // big buffer
			tmp := make([]byte, 256) // using small tmo buffer for demonstrating
			for {
				n, err := conn.Read(tmp)
				if err != nil {
					if err != io.EOF {
					}
					break
				}
				all = append(all, tmp[:n]...)
			}
			fmt.Printf(string(all[:]))
		}
	}
}
