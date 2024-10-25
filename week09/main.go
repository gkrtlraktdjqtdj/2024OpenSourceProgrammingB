package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // dice 1~6
	fmt.Println(answer)

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다 숫자 입력 : ", 3-guesses)
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input) // Atoi : sting 을 integer 로 바꿈
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			fmt.Println("정답입니다!")
			break
		} else if answer > guess {
			fmt.Println("입력하신 값은 작습니다")
		} else {
			fmt.Println("입력하신 값이 큽니다")
		}
	}
}
