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

	alphaGroup1 := [3]string{"Damenjo", "Tari", "Fadhil"}
	fmt.Print(alphaGroup1)
	alphaGroup2 := [...]string{"PHP", "Phyton", "Golang", "C++", "Java"}
	fmt.Print(alphaGroup2)
}
