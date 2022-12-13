package main

import "fmt"

func main() {
	// 실수 자료형을 사용할 수 있다.
	// float32(7자리), float64(15자리)

	// 예제1
	var num1 float32 = 0.14
	var num2 float32 = .75647 // 소수점 앞에 0 하나는 생략할 수 있다.
	var num3 float32 = 332.0378373
	var num4 float32 = 10.0

	var num5 float32 = 14e6
	var num6 float32 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println("num1 : ", num1)
	fmt.Println("num2 : ", num2)
	fmt.Println("num3 : ", num3)

	fmt.Println("num4 : ", num4)
	fmt.Println("num4 : ", num4-0.1)
	fmt.Println("num4 : ", float32(num4-0.1))
	fmt.Println("num4 : ", float64(num4-0.1)) // 주의! 자료형 대역폭이 넓어졌기 때문에 부동소수점 오차가 발생할 수 있다.

	fmt.Println("num5 : ", num5)
	fmt.Println("num6 : ", num6)
	fmt.Println("num7 : ", num7)
}
