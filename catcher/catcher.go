package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(0)
	}
}

func MessageCatcher(conn *net.UDPConn) {
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		log.Printf("%d Received: %s from %s", n, string(buf[0:n]), addr)

		if err != nil {
			log.Println("Error: ", err)
		}
	}
}



func main() {

	serverAddr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	// Listen at selected port
	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)

	defer serverConn.Close()

	go MessageCatcher(serverConn)

	fmt.Scanln()
}