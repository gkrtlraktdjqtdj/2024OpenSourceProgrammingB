package greeting

import "fmt" //package 이름은 모두 소문자고 카멜 표기법이 적용되지 않는다

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
