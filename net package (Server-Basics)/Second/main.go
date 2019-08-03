package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	go dialingClient()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

// Run using `telnet localhost 8080`

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "Oh hey! Here's you msg: %s", ln)
	}
	defer conn.Close()
}

func dialingClient() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		fmt.Fprintln(conn, "Wew, hey!")
	}
	fmt.Println("Pinged the server 3 times with 'Wew, hey!'")
}
