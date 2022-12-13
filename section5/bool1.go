package main

import "fmt"

func main() {
	// Bool(Boolean) : 참과 거짓
	// 조건부 논리 연산자(!, ||, &&)와 함께 주로 사용된다.
	// 유의사항 : 암묵적 형 변환이 안된다.

	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("b1 : ", b1)
	fmt.Println("b2 : ", b2)
	fmt.Println("b3 : ", b3)

	fmt.Println("b3 == b3 : ", b3 == b3)
	fmt.Println("b1 && b3 : ", b1 && b3)
	fmt.Println("b1 || b2 : ", b1 || b2)
	fmt.Println("!b1 : ", !b1)
}
