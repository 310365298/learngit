package main

import (
	"fmt"
	"time"
)

func main(){
	a := time.Now().Unix()
	switch{
	case (a - 64800) % 8640 == 0 :fmt.Printf("晚安，打工人")
	case a % 8640 == 0 : fmt.Printf("早安，打工人")
	}
}
