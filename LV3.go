package main

import "fmt"

var ch3 = make(chan int)

func main() {

	over := make(chan bool,1)	//将over改为缓存管道

	for i := 0; i < 10; i++ {

		go func() {

			ch3 <- i

		}()
		res2 := <- ch3
		fmt.Printf("%v\n", res2)

	if i == 9 {

	over <- true

		}
	}

<-over

fmt.Println("over!!!")

}