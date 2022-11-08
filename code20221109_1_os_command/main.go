package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("hello world")

	if len(os.Args) == 1 {
		fmt.Println("매개변수 더 입력하시오")
		os.Exit(1)
	}

	args := os.Args
	//두번 째 값 _ 에러는 무시
	min, _ := strconv.ParseFloat(args[1], 64)
}
