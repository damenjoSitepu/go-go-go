package main

import (
	"fmt"
)

func main() {
	fmt.Println("Basic Array")
	var myNumber [2]int
	myNumber[0] = 100
	firstIndex := myNumber[1]
	fmt.Println(firstIndex)
}
