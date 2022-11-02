package main

import "fmt"

func notGenericMinCalc(num0001 int, num0002 int) int {
	if num0001 < num0002 {
		return num0001
	} else {
		return num0002
	}
}

type IntegerList interface {
	int | int8 | int16 | int32 | int64
}

type FlotList interface {
	float32 | float64
}

type NumberList interface {
	FlotList | IntegerList
}

// [T any] 타입 파라미터 : 함수에서 사용할 제네릭 파라미터 선언 ~타언어 <> ~
// any 모든 타입 가능
// useGenericMinCalc(a interface{}) 모든 타입은 인터페이스 가능하나 operation 불가능
// useGenericMinCalc[T int | int16] or절 대신에 type interface 활용
func aaa[T IntegerList](input T) T {
	return input
}

func useGenericMinCalc[T NumberList](a, b T) T {
	if a < b {
		return a
	}
	return b
}
func main() {
	fmt.Println("hello world")
	var num1 int16 = 10
	var num2 int16 = 20
	//fmt.Println(notGenericMinCalc(num1, num2))
	fmt.Println(useGenericMinCalc(num1, num2))

}
