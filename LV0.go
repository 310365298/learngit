package main

import (
	"fmt"
)

var m = make(map[int]int,10)
var ch = make(chan int)
func jiecheng(i int){
	res := 1
	for x := 1; x <= i; x++ {
		res *= x
	}
	ch <- res
}

func main() {
	for i := 1; i <= 20; i++{
		go jiecheng(i)
		m[i] = <- ch
	}
	for m, v := range m {
		fmt.Printf("m[%v]=%v\n", m, v)
	}
}