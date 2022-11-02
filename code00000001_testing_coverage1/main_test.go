package code00000001_testing_coverage1

import (
	"testing"
)

// fileName은 *_test.go
// func 이름은 Test로 시작
// 매개변수는 t *test.T
// 실패지점 t.Fail() / t.Error 호출

// go test -cover / go test -v
// go test [PACKAGE_NAME]

// panic 테스트
// recover가 panic보다 위에 있어야 함

// 학습 테스트
func TestAppend(t *testing.T) {
	sliced := []int{1, 2, 3, 4, 5}
	sliced = append(sliced, 6)

}

// report
// go test -coverprofile=coverage.out / go tool cover -html=coverage.out
// go test -bench=. : 속도
// go test -bench=. -benchmem : 속도, 메모리
// go test -bench .

func TestCheck(t *testing.T) {
	v := testFunc0002()
	if v != 4 {
		t.Error("Fails!!!")
	}
	t.Log("-v flag")
}

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibo(50)
	}
}

func BenchmarkFibo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFibo2(50)
	}
}
