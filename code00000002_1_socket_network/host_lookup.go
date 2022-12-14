package main

import (
	"fmt"
	"net"
	"os"
)

// go run [FILE_NAME] www.google.co.kr
// domain ip주소 찾기
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	// domain 매개변수 받아서 사용하는 ip 목록 출력
	addrs, err := net.LookupHost(name)
	// net.ResolveIPAddr("ip", name) // ip 찾기
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
