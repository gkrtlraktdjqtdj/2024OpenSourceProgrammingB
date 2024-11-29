package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)

	var name string
	var age int
	for {
		fmt.Print("Name? ")
		fmt.Scanln(&name)
		if name == "q" {
			break
		}
		fmt.Print("Age? ")
		fmt.Scanln(&age)

		ages[name] = age
	}
	for key, value := range ages {
		fmt.Printf("%s is %d years old\n", key, value) //value = key[ages] or ages[key]
	}
	//슬라이스는 배열과 달리 한타입만 넣을 수 있다
	//create map (like) python dict
	//go map은 순서 개념이 없다
}
