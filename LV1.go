package main

import "fmt"

var (
	ch1 = make(chan int)
	ch2 = make(chan int)
	i int
	res1 int
	res2 int
	key int
)

func gor1(i int){
	for i = 0; i <= 101; i += 2{
		ch1 <- i
	}
}
func gor2(i int){
	for i = 1; i <= 100; i += 2{
		ch2 <- i
	}
}
func main(){
	go gor1(i)
	go gor2(i)
	for key = 0; key<=50; key++{
		res1 = <- ch1
		res2 = <- ch2
		fmt.Printf("%v %v ",res1,res2)
	}
}
