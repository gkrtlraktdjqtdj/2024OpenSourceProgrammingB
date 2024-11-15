package main

import (
	"fmt"
	"time"
)

func main() {
	//var dates [3]time.Time
	//dates[0] = time.Unix(0, 0)
	//dates[2] = time.Unix(1708012346, 0)
	//fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	//var dates [3]time.Time = [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	//fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	//dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	//fmt.Println(dates[0], dates[1], dates[2]) // unix time, zero value, 2024-02-16 ...

	//dates := [3]time.Time{
	//	time.Unix(0, 0),
	//	time.Unix(1, 0),
	//	time.Unix(1708012346, 0), // , 여러줄일 때 필수
	//}
	//fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012346, 0)} // 중괄호를 내리지 않으면 가능
	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Println(dates)         // array
	fmt.Printf("%#v\n", dates) // array litaral

}
