package main

import (
	"GameSocketServer/pkg"
	"log"
	"net"
)

func main() {
	s := pkg.NewServer()

	go s.Run()

	listener, err := net.Listen("tcp", "127.0.0.1:8888")

	if err != nil {
		log.Fatalf("Unable to start server: %v", err.Error())
	}

	defer listener.Close()
	log.Printf("starter server on port : 8888")

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("starter to accepr connection: %v", err.Error())
			continue
		}
		go s.NewClient(conn)
	}
}
