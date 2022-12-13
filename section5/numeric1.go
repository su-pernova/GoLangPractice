package main

import "fmt"

func main() {
	// 숫자형 데이터 타입(정수, 실수, 복소수)을 사용할 수 있다.
	// 32 bit, 64 bit, unsigned(양수)
	// 정수 : 8진수(0), 16진수(0x), 10진수
	// Golang 의 특징 : 자료형을 엄격하게 지킨다. 효과 :  가독성이 좋다.

	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)
	fmt.Println("num4 : ", num4)

}
