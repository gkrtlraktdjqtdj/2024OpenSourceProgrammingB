package main

import "fmt"

func main() {
	var student struct {
		id   int
		name string
		gpa  float32
	}

	student.id = 202444051
	student.name = "son heungmin"
	student.gpa = 4.5
	fmt.Println(student.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 202444052
	student2.name = "seo jihye"
	student2.gpa = 4.4
	fmt.Println(student2.gpa)
	// 슬라이스 ,멥은 한가지 타입만 할당받을 수 있다
	// map[string]int map[key type]value type
	// %#v = reflet.typrOf()
}
