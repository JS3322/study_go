package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server on!")
	http.ListenAndServe(":3000", nil)

}

func checkNilErr3(err error) {
	if err != nil {
		panic(err)
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Method : ", r.Method)
	log.Println("URL : ", r.URL)
	log.Println("Header : ", r.Header)

	response, err := io.ReadAll(r.Body)
	checkNilErr3(err)
	defer r.Body.Close()
	fmt.Println("Body : ", string(response))

	switch r.Method {
	case "POST":
		w.Write([]byte("POST request success"))
	case "GET":
		w.Write([]byte("GET request success"))

	}

	//text/plain : html 안먹힘
	//text/html : tag 먹힘
	w.Header().Set("Content-Type", "text/plain")
	//w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello world!</h1>")
}
