package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(0)
	}
}

func MessageSender(src *net.UDPAddr, dst *net.UDPAddr) {

	Conn, err := net.DialUDP("udp", src, dst)
	CheckError(err)

	defer Conn.Close()

	i := 0
	for {
		msg := fmt.Sprintf("Hi - %s", strconv.Itoa(i))
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 2)
	}
}



func main() {

	//srcAddr, err := net.ResolveUDPAddr("udp", fmt.Sprint("127.0.0.1", ":10002"))
	//CheckError(err)
	desAddr, err := net.ResolveUDPAddr("udp", fmt.Sprint("127.0.0.1", ":10001"))
	CheckError(err)


	go MessageSender(nil, desAddr)
	//// Create connection
	//Conn, err := net.DialUDP("udp", srcAddr, desAddr)
	//CheckError(err)
	//
	//defer Conn.Close()
	//
	//// Create message
	//msg := fmt.Sprintf("Hi (from %s)", "xxxxx")
	//
	//buf := []byte(msg)
	//
	//n, err := Conn.Write(buf)
	//if err != nil {
	//	log.Printf("Error message: %s, (%s(%d))\n", err, msg, n)
	//}
	fmt.Scanln()

}
