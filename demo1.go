package main

import "fmt"

func main(){
	//var m1 map[string]int = make(map[string]int)
	//m1["a"] = 1
	//m1["b"] = 2
	//m1["c"] = 3
	m1 := map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
	fmt.Println(m1["a"])
	fmt.Println(m1)
}
