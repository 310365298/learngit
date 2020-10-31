package main

import (
	"fmt"
)

func Receiver(v interface{}) {
	switch v.(type){
	case int:fmt.Println("这个是int")
	break
	case string:fmt.Println("这个是string")
	break
	case bool:fmt.Println("这个是bool")
	break
	}
}

func main(){
	Receiver("你好吗")
	Receiver(32)
	Receiver(true)
}
