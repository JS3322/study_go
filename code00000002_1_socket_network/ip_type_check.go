// https://www.joinc.co.kr/w/man/12/golang/networkProgramming/SocketProgramming
// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/ip_address_type.html
// https://organicprogrammer.com/2021/07/31/how-to-implement-simple-http-server-golang/
package main

//net
//type IP []byte
//type IPMask []byte
//func IPv4Mask(a, b, c, d byte) IPMask
//func (ip IP) DefaultMask() IPMask
import (
	"fmt"
	"net"
	"os"
)

// go run [FILE_NAME] 127.0.0.1
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	//IPv4와 IPv6 처리하는 함수
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	} else {
		fmt.Println("THe address is ", addr.String())
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	/*
		주소는 127.0.0.1로 클래스 주소다.
		default mask의 크기는 32bit다.
		A클래스이니 앞 8비트를 이용해서 마스크 연산을 한다.
		마스크는 ff000000(255.0.0.0)이다.
		네트워크는 127.0.0.0 이다.
	*/
	fmt.Println("Address is ", addr.String(),
		"\nDefault mask length is ", bits,
		"\nLeading ones count is", ones,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
	os.Exit(0)
}
