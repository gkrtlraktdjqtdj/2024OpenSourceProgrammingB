package greeting

import "fmt"

func hi(name string) { // 함수 이름 첫글자는 대문자여야함
	fmt.Printf("HI %s!\n", name) // 소문자로 시작하는 함수의 경우 접근 불가
} //단 내부에서는 접근 가능
func hello(name string) {
	fmt.Printf("Hello %s!\n", name)
}
func EnglishGreetings(name string) { // 공개된 함수를 통해 접근
	hello(name)
	hi(name)
}
