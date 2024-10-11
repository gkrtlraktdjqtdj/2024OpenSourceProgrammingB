package main

import (
	"bufio" //Buffer input output
	"fmt"
	"os" // Operating System
)

func main() {
	fmt.Print("이름 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') //RS는 두 값를 리턴함 ,뒤에 _를 쓰면 nil값이 안보이게됨
	fmt.Println(i)
	fmt.Println(err) //에러가 없을 때 <nil>값을 줌

}
