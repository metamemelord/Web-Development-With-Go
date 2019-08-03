package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "Hello world!!")
		fmt.Fprintln(conn, "Hello world!!(2)")
		fmt.Fprintf(conn, "Hello world!!(2)\n")

		conn.Close()
	}

}
