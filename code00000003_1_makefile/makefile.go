package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "this is cleancode text"

	file, err := os.Create("./testFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is :", length)

	readFile("./testFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("text data is \n", databyte)
	fmt.Println("string is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
