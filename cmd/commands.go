package main

import (
	"fmt"
	"net"
)

func main(){
	address := ":6379"

	ln, err := net.Listen("tcp", addr)
	
	fmt.Println("listening on", addr)

	
	for {
		
		conn, err := ln.Accept()
		if err != nil {
			
			fmt.Println("accept error:", err)
			continue
		}

		
		go handleConn(conn)
	}
}
