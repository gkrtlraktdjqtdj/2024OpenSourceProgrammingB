package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	r := fmt.Sprintf("%d\n", rand.Intn(6)+1) // Sprintf : Printf와 같은 동작, 문자열을 리턴, 출력 X
	fmt.Println(reflect.TypeOf(r), reflect.TypeOf(rand.Intn(6)+1))
	fmt.Printf("%T\n", r)

	i := 1
	for i <= 10 {
		fmt.Printf("%2d점\n", i)
		i = i + 1
	}
}
