package main

import (
	"fmt"
	"go-go-go/print"
	"log"
)

func main() {
	log.SetPrefix("Hina: ")
	log.SetFlags(0)
	// message := print.Hina("Tari Puspita Sari")
	// message, e := print.Hina("Damenjo Sitepu")
	message, e := print.Hina("")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(message)
}
