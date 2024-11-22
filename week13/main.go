package main

import (
	"fmt"
)

func main() {

	gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array litaral
	gpa_slice := gpas[1:4]                       // [3.5, 4.1, 3.9]
	//gpa_slice[1] = 2.71
	gpas[2] = 2.71
	//gpa_slice = append(gpa_slice, 4.3)
	gpa_slice = append(gpa_slice, 4.3, 5.55)
	fmt.Println(len(gpa_slice), gpa_slice, gpas) // 배열 크기의 관계가 배열 >= 슬라이스일 땐 한쪽을 바꾸면 둘 다바뀌지만 슬라이스의 크기가 배열의 크기보다 커지면 슬라이스가 따로 저장공간을 할당받는다
}
