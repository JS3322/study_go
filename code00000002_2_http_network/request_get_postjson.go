package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("hello world")
	//PerformGetRequest()
	PerformPostJsonRequest()
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest() {
	const url_get_source = "https://dummyjson.com/todos/1"

	response, err := http.Get(url_get_source)
	checkNilErr(err)

	defer response.Body.Close()
	fmt.Println("status code : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	//fmt.Println(content)
	fmt.Println(string(content))

	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is : ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {

	const myrul = "https://dummyjson.com/todos/1"

	requestBody := strings.NewReader(`
		{
			"coursename" : "golang start",
			"price": 0,
			"platform" : "cleancode.kr"
		}
	`)
	response, err := http.Post(myrul, "application/json", requestBody)

	checkNilErr(err)
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	log.Println(string(content))
}
