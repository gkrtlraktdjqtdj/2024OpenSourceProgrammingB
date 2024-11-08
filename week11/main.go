package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
	//"week11/week11/keyboard"
)

func main() {
	//n, _ := keyboard.GetFloat()
	n, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}
