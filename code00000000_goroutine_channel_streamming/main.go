package main

import "fmt"

// https://judo0179.tistory.com/88
// 물리 CPU 개수만큼 사용하도록 설정되어 각 코어에서 시분할 ( Concurrent ) 처리로 동작합니다. 만약 여러 CPU를 병렬로 실행하고자 하는 코드를 작성하고자 하면 runtime.GOMAXPROCS(runtime.NumCPU()) 함수를 호출하여 시스템의 모든 코어를 사용하도록 설정
// 채널은 참조타입
func main() {
	//chain 버퍼크기(타입)
	c := make(chan int)
	go sum(10, 20, c)

	result := <-c
	fmt.Println(result)
}
func sum(a int, b int, c chan int) {
	c <- a + b
}

//채널에는 Unbuffered Channel과 Buffered Channel 2가지 채널이 존재
