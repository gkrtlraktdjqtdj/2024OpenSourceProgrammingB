package main

import (
	"bufio" //Buffer input output
	"fmt"
	"log"
	"os" // Operating System
)

func main() {
	fmt.Print("이름 입력 : ")
	r := bufio.NewReader(os.Stdin)
	name, err := r.ReadString('\n') //RS는 두 값를 리턴함 ,뒤에 _를 쓰면 nil값이 안보이게됨

	if err != nil { //에러가 났을 경우 프로그램 종료
		log.Fatal(err)
	} else { //go는 c처럼 중괄호로 꼭 묶어줘야 함
		fmt.Println(name)
	}

}
