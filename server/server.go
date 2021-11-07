package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1200")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
	select{}
}
func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(1 * time.Minute)) // set 2 minutes timeout
	log.Println("adress:", conn.LocalAddr().String())
	defer conn.Close() // close connection before exit
	for {
		conn.Write([]byte("123"))
		request := make([]byte, 128)
		readLen, err := conn.Read(request)
		//request, err := ioutil.ReadAll(conn)
		if readLen == 0 {
			break // connection already closed by client
		}
		log.Println(string(request))
		if err != nil {
			break
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

