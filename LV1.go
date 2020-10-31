package main

import "fmt"

//var(
	a string
	b string
	c string
)

//type Author struct {
	VIP bool
	Icon string
	Signature string
	Focus int
}
//type Video struct {
	Name string
	Hits int
	Praised int
	Coin int
}
	func main(){
	var xiaochao Author
	var jiuzhe Video
	xiaochao.VIP=true
	xiaochao.Icon="xiaochaoyuanzhang"
	xiaochao.Signature="小潮team一鸽"
	xiaochao.Focus=4637000
	jiuzhe.Name="jiuzhe"
	jiuzhe.Hits=2600000
	jiuzhe.Praised=432000
	jiuzhe.Coin=366000
	fmt.Println(jiuzhe)
	fmt.Println(xiaochao)
}
