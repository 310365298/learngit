package main

import (
	"io"
	"os"
)

func main() {
	a,err := os.Create("./proverb.txt")
	a,err = io.Writer("don't communicate by sharing memory share memory by communicating")
}
