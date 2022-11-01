package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//*(python)dic, (javascript)object
	//map[key]value
	mapExample := make(map[string]int)
	mapExample["port"] = 9090
	fmt.Println(mapExample["port"])
}
