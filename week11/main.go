package main

import (
	"fmt"
	"log"
	"week11/week11/keyboard"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		i := 3
		for i*i <= n {
			if n%i == 0 {
				return false
			}

			i = i + 2
		}
	}
	return true
}
func main() {
	fmt.Print("첫번째(시작) 정수 입력 : ")
	n1, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("두번째(끝값) 정수 입력 : ")
	n2, err := keyboard.GetInteger()
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
