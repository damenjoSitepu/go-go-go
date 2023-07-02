package main

import (
	"fmt"
	// "go-go-go/print"
	"go-go-go/grade"
	"log"
)

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)
	// message := print.Hina("Tari Puspita Sari")
	// message, e := print.Hina("Damenjo Sitepu")
	// message, e := print.Hina("aso")
	// message, e := print.PrintRandomMessage("My Dear")
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// fmt.Println(message)
	result, e := grade.PassOrNot("B")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(result)
}
