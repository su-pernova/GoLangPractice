package main

import "fmt"

func main() {
	// 상수(const)를 선언하고 초기화할 수 있다.
	// 반드시 선언과 초기화가 동시에 이루어져야 하며 한 번 할당한 값을 변경할 수 없다.
	// 고정된 값을 사용하고 관리할 때 유용하다.

	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	const d = 35.6
	const e = false

	/*
		< 에러가 발생하는 케이스 >
		1. 선언과 초기화가 동시에 이루어지지 않을 경우
			const g string
			g = "Test3"
		2. 함수의 return 값은 상수에 할당할 수 없다.
			const g = getHeight()
	*/

	// 변수와 달리 선언하고 사용하지 않아도 에러가 발생하지 않는다.
	fmt.Println("a : ", a, " / b : ", b, " / c : ", c, " / d : ", d, " / e : ", e)

}
