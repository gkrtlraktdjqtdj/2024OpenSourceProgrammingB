package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
// func test(strs ...string, i int) { //error
func test(i int, strs ...string) { //가변매게변수는 항상 맨뒤에 있어야 함
	fmt.Println(i, strs, reflect.TypeOf(strs))
}

func main() {
	//fmt.Println(os.Args, len(os.Args), reflect.TypeOf(os.Args)) // go run main.go I Love You // ./main.exe U love me 리눅스 스타일
	slices := os.Args[1:] //커멘드창에 직접 입력
	fmt.Println(slices, slices[2])
	test(1, "123")
	test(-99, "123", "abc")
	test(20)
	test(45, "123", "abc", "123a")
}
