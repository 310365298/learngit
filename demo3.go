package main

import "fmt"

func main(){
	defer func() {
		panic("first")
	}()
	defer func() {
		panic("second")
	}()
	fmt.Println("main function")
}
