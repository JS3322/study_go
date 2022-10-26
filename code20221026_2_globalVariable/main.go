package main

import (
	"fmt"
)

// package global variable
var globalVariable int = 55

func main() {
	fmt.Println("hello world")

	var globalVariable int = 20
	//가까운 main메서드의 지역변수를 호출
	fmt.Println(globalVariable)

	{
		var localVariable int = 21
		var globalVariable int = 12
		//가까운 지역변수를 호출
		fmt.Println(globalVariable)
		fmt.Println(localVariable)
	}
	//지역변수 호출 불가
	//fmt.Println(localVariable)

	//bool
	//string (builder, buffer)
	//array
	//slice
	//struct
	//point
	//func
	//map
	//interface
	//chanel

	//not use builtin
	//similar to scan
	fmt.Print("not newline")
	fmt.Println("newline")
	//순서대로 서식문자에 대입
	fmt.Printf("%d or %f\n", 20, 3.232)
	//%d 10진수, %b 2진수, %T 데이터타입
}
