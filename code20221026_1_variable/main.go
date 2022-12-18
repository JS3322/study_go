package main

import (
	"fmt"
)

func main() {
	//var 변수선언 | num 변수명 | int 변수타입
	// = 대입연산자
	var num int = 30
	//: var
	number := 90
	var testNum int32 = 20
	print(testNum)
	var message string = "hello world"
	aSlice := []int{1, 2, 3, 4, 5}
	//fmt 패키지의 Println : 에러가 날 수 있다
	//func Println(a ...any) (n int, err error) {
	//	return Fprintln(os.Stdout, a...)
	//}
	fmt.Println(aSlice)
	fmt.Println(message, num, number)

	fmt.Println("-------------------1")

	//unit8=byte 16 32 64 (1,2,4,8 바이트) 부호없는 정수
	//int8 16 32=int(32bit), 64(64bit) (1,2,4,8 바이트) 부호있는 정수 // 32 = rune
	//float32 64 (4, 8 바이트) 실수
	//complex64 128(진수, 가수) 복소수,

	//int8의 1바이트 수치 제외하고 버려짐
	var num0001 int32 = 325235
	var num0002 int8 = int8(num0001)
	fmt.Println(num0001, num0002)
}
