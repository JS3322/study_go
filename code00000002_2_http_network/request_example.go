package main

import (
	"fmt"
	"io"
	"net/http"
)

const url_code = "https://log.cleancode.kr"

func main() {
	fmt.Println("cleancode webPage request")
	response, err := http.Get(url_code)
	if err != nil {
		panic(err)
	}
	fmt.Println("response is of type : \n", response)

	defer response.Body.Close() // close the connection

	databytes, err := io.ReadAll(response.Body)
	//io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
