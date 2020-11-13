package main

import (
	"fmt"
	"strings"
)

func main(){
	str := "hello world"
	fmt.Println(strings.Index(str, "e"))
	fmt.Println(strings.Index(str, "c"))
	fmt.Println(strings.IndexAny(str, "o"))
}
