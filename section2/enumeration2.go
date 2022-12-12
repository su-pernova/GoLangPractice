package main

import "fmt"

func main() {
	// iota 라는 예약어를 이용해 열거형을 사용할 수 있다.
	// 값을 일일이 할당해 줄 필요 없이, 처음 값을 통해 모든 값 초기화 가능
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println("Jan : ", Jan)
	fmt.Println("Feb : ", Feb)
	fmt.Println("Mar : ", Mar)
	fmt.Println("Apr : ", Apr)
	fmt.Println("May : ", May)
	fmt.Println("Jun : ", Jun)
}
