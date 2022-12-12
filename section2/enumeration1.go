package main

import "fmt"

func main() {
	// 열거형 : 일정한 규칙에 따라 숫자가 계산되거나 증가하는 상수의 묶음
	const (
		Jan = 1
		Feb = 2
		Mar = 3
		Apr = 4
		May = 5
		Jun = 6
	)

	fmt.Println("Jan : ", Jan)
	fmt.Println("Feb : ", Feb)
	fmt.Println("Mar : ", Mar)
	fmt.Println("Apr : ", Apr)
	fmt.Println("May : ", May)
	fmt.Println("Jun : ", Jun)
}
