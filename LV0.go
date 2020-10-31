package main

import "fmt"
type yuzige struct {
	who string
	why string
	how string
}
func main(){
	var a yuzige
	a.who="男神"
	a.why="又帅又有才"
	a.how="非常非常非常男神"
	fmt.Println(a)
}
