package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	std := bufio.NewReader(os.Stdin)

	fmt.Println("hello world")

	numExample1 := 11
	numExample2 := &numExample1
	var numExample3 *int = &numExample1
	fmt.Println(numExample2)  // numExample1 의 주소
	fmt.Println(*numExample2) // '' 의 값
	fmt.Println(numExample3)
	fmt.Println(*numExample3)

	//&메모리 주소

	var num0001 int
	//num0002 := 20
	var num0002 int = 50

	//buffer
	//data 한 번에 받아서 연결
	n, err := fmt.Scanln(&num0001, &num0002)
	if err != nil {
		fmt.Println("err!")
		fmt.Println(err)
		std.ReadString('\n')
	} else {
		fmt.Println(n, num0001, num0002)
	}

}
