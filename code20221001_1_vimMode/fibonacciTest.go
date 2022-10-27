package main

import (
	"fmt"
	"time"
)

func vimTest() {
	//hjkl 상하좌우
	//x 삭제 dd줄삭제
	//u 되돌리기 p 되돌리기(커서 위치에서)
	//r 치환(대신 바꿀 글자)
	//숫자 + shift + G 줄 위치로 이동
	//ctrl + g 단어블록
	//? 윗방향으로 검색 / 아랫방향으로 검색 / 검색 후 n 다음검색어 shift + n 반대방향 다음검색어
	// o 위로 한 줄 열고 커서 모드 a 아래로 한 줄 열고 커서 모드
	fmt.Println("hallo world")

}

func Fib(max uint64) <-chan uint64 {
	ch := make(chan uint64)
	go func() {
		defer close(ch)
		var a, b uint64
		a, b = 0, 1
		for a <= max {
			ch <- a
			a, b = b, a+b
		}
	}()
	return ch
}

func FibGenerator(max uint64) func() uint64 {
	var next, a, b uint64
	next, a, b = 0, 0, 1
	return func() uint64 {
		next, a, b = a, b, a+b
		if next > max {
			return 0
		}
		return next
	}
}

func main() {
	start := time.Now()
	for fib := range Fib(9999999999999999999) {
		fmt.Print(fib, ", ")
	}
	end := time.Since(start)
	fmt.Println("END!", end)
}
