package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[0] = time.Unix(0, 0)
	// dates[2] = time.Unix(1708012346, 0)
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	// var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	// dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	// fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	// dates := [3]time.Time{
	//	 time.Unix(0, 0),
	//	 time.Unix(1, 0),
	//	 time.Unix(1708012346, 0), // comma , 여러줄일 때 필수
	// }
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012346, 0)} // 중괄호를 내리지 않으면 가능
	// fmt.Println(dates[0], dates[1], dates[2])
	// fmt.Println(dates)         // array
	// fmt.Printf("%#v\n", dates) // array litaral

	// for i := 0; i <= 7; i++ {
	// for i := 0; i < len(dates); i++ { // =을 뺴야 함
	// 	fmt.Println(i, dates[i])
	// }

	for i, date := range dates { // index 사용
		fmt.Println(i, date)
	}

	// for _, date := range dates { // 인덱스 미사용 like python for in, SAFE
	// 	fmt.Println(date)
	// }

}
