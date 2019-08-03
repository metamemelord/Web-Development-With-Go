package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
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
		go handle(conn)
	}
}

// Run using `telnet localhost 8080`
// > hehe
// urur

func handle(conn net.Conn) {
	fmt.Println("Opening connection for 60 seconds...")
	err := conn.SetDeadline(time.Now().Add(1 * time.Minute))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := caesar13(bs)
		fmt.Fprintf(conn, string(r)+"\n")
	}
	defer conn.Close()
}

func caesar13(bs []byte) []byte {
	var result = make([]byte, len(bs))
	for i, v := range bs {
		if v <= 109 {
			result[i] = v + 13
		} else {
			result[i] = v - 13
		}
	}
	return result
}
