package main

import (
	"fmt"

	"github.com/biancarosa/algo/unionfind"
)

func main() {
	fmt.Println("hello")
	quickfind := unionfind.QuickFindImpl{}
	quickfind.New(3)
	fmt.Println(quickfind.Connected(0, 1))
	fmt.Println(quickfind.Connected(1, 2))
	fmt.Println(quickfind.Connected(0, 2))
	quickfind.Union(1, 2)
	fmt.Println(quickfind.Connected(0, 1))
	fmt.Println(quickfind.Connected(1, 2))
	fmt.Println(quickfind.Connected(0, 2))
	quickfind.Union(0, 2)
	fmt.Println(quickfind.Connected(0, 1))
	fmt.Println(quickfind.Connected(1, 2))
	fmt.Println(quickfind.Connected(0, 2))
}
