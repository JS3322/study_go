package main

import (
	"fmt"
	"time"
)

func main() {
	num := 20
	if num <= 20 {
		fmt.Println("num 20 이상")
	}
	for i := 1; i < 100; i++ {
		fmt.Println("hello")

	}

	go func0001()
	func0002("world")
}

// func 함수명(매개변수) 리턴타입
func func0001() int {
	for i := 1; i < 100; i++ {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}
	return 20000
}

func func0002(input string) {
	for i := 1; i < 100; i++ {
		fmt.Println(input)
		time.Sleep(time.Second)
	}

}
