package main

import "github.com/common-nighthawk/go-figure"

// memory buffer
var buf [1024]byte

//export bufAddr
func bufAddr() *byte {
	return &buf[0]
}

//export test
func test(s string) {
	myFigure := figure.NewFigure(s, "", true)
	myFigure.Print()
}

func main() {}
