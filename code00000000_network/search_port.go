package main

import (
	"fmt"
	"net"
	"os"
)

// go run [FILE_NAME] tcp ssh
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage : %s network-type service\n",
			os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	service := os.Args[2]
	// tcp udp 방식으로 어떤 서비스가 몇 번 포트를 사용하는지 확인
	port, err := net.LookupPort(networkType, service)
	//ResolveTCPAddr 도메인과 서비스타입으로 아이피 포트 반환
	tcpAddr, err := net.ResolveTCPAddr("tcp", "www.google.co.kr:ssh")
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

	fmt.Println("Service Port", port)
	fmt.Println("IP ", tcpAddr.IP.String(), "Port ", tcpAddr.Port)
	os.Exit(0)
}
