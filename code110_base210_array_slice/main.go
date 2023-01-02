package main

import "fmt"

func main() {
	fmt.Println("hello world")

	//array
	arrayExample := [3]int{20, 10, 30}
	fmt.Println(arrayExample)

	//slice *(java)arraylist
	sliceExample := []int{}
	sliceExample = append(sliceExample, 90, 30, 40)
	fmt.Println(sliceExample)
}
