package main

import (
"fmt"
"net"
"os"
"time"
)
func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1200")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn.Write([]byte("fi:1234"))
	checkError(err)
	for {
		result := make([]byte, 256)
		_, err = conn.Read(result)
		checkError(err)
		fmt.Println("get mes:", string(result))
		time.Sleep(time.Second * 1)
		conn.Write([]byte("this is fi" + time.Now().Format("2006-01-02 15:04:05")))
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}


