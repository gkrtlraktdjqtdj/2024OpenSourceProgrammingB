package main

import (
	"fmt" // c language #include <stdio.h>
	"math"
	"reflect"
	"strings"
)

func main() {
	// var i int
	// i = 55

	//var i int = 55

	var f float64 = 1.92

	// f := 1.92 //기본으로 64비트 할당
	i := "55"
	fmt.Println(strings.Title("kim inha")) //단어앞글자 대문자로 바꿈
	fmt.Printf("%f %f\n", f, math.Ceil(f)) //math.floor(내림) math.Ceil(올림)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("f is", f)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n")
	fmt.Println("i is", i)
	//fmt.Printf("i is %d\n", i)
	fmt.Printf("i is %s\n", i)
}
