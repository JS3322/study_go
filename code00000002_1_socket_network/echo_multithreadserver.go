package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// ps -eLf | grep main
// cd /proc/3077/fd
func main() {
	service := "0.0.0.0:1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// go routine 실행
		go handleclient(conn)
	}
}

func handleclient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	fmt.Println("PID : ", os.Getpid())
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
