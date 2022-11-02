package code00000001_testing_coverage2

import "testing"

func Test_calc(t *testing.T) {
	//result
	result := 5.0
	//test func
	v := clac(5)
	if v != result {
		t.Error("Fails!!!")
	}
	t.Log("-v flag")
}

func Benchmark_clac(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clac(20)
	}
}
