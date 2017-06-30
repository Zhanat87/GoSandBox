package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

var clients = []net.Conn{}

func serve() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	log.Println("Listening on port 8000...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	clients = append(clients, conn)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		for _, c := range clients {
			c.Write([]byte("\t" + input.Text() + "\n"))
		}
	}
	conn.Close()
}

func conn() {
	c, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Println("Failed to connect: ", err)
	}
	go io.Copy(c, os.Stdin)
	io.Copy(os.Stdout, c)
}

func main() {
	switch os.Args[1] {
	case "serve":
		serve()
	case "conn":
		conn()
	default:
		log.Println("Broken")
	}
}
