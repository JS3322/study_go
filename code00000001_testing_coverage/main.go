package code00000001_testing_coverage

func testFunc0001() int {
	return 2
}

func testFunc0002() int {
	return 4
}

func getFibo(n int) int {
	n1, n2 := 1, 2

	for i := 0; i < n-2; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2

}

func getFibo2(n int) int {
	var slice []int
	slice = []int{1, 1}

	for i := 0; i < n-2; i++ {
		slice = append(slice, slice[i]+slice[i+1])
	}

	return slice[len(slice)-1]
}

func main() {

}
