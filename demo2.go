package main

import "fmt"

func main(){
	var(
		a int
		b int
		c int
	)
	fmt.Scanf("%d",&a)
	defer fmt.Printf("%d\n", c) //c=0
	fmt.Scanf("%d",&b)
	c = a + b
	defer fmt.Printf("%d\n", c) //c=a+b
}
