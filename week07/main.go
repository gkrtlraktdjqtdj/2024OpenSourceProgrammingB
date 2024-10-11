package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		score = strings.TrimSpace(score)                //줄바꿈, 띄어쓰기 ,탭 등 제거 파이썬의 strip과 유사
		realScore, _ := strconv.ParseInt(score, 16, 32) // 정수형 32비트 타입으로 형변환

		if realScore >= 60 {
			fmt.Println("A")
			fmt.Println("%d", realScore)
		} else {
			fmt.Println("BCDF")
			fmt.Println("%d", realScore)
		}
	}
}
