package main

import "fmt"

func main() {
	// 형변환 시에는 조금 더 큰 범위의 자료형 쪽으로 맞춰주는 것이 좋다.
	// 예) 정수 + 실수 -> 정수를 실수로 형변환
	// 형변환 시에는 해당 자료형이 표현가능한 min, max 값도 반드시 유의하자.

	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println(n1 + n2)
	fmt.Println(n1 - n2)
	fmt.Println(n1 * n2)
	fmt.Println(n1 / n2)
	fmt.Println(n1 % n2)
	fmt.Println(n1 << n2)
	fmt.Println(n1 >> n2)
	fmt.Println(^n1)

}
