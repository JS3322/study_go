package main

import (
	"fmt"
	"net/url"
)

const url_source string = "https://sports.news.naver.com/news?oid=449&aid=0000237570"

func main() {
	fmt.Println("hello world")
	fmt.Println(url_source)
	//parsing
	result, _ := url.Parse(url_source)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	fmt.Println("------------------")
	qparams := result.Query()
	fmt.Println("query params are : \n", qparams)
	fmt.Println(qparams["oid"])

	fmt.Println("------------------")
	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	fmt.Println("------------------")
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "naver.com",
		Path:    "/news",
		RawPath: "user=js3322",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
