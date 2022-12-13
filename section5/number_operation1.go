package main

import (
	"fmt"
	"math"
)

func main() {
	// 숫자 자료형의 연산을 할 수 있다.
	// 타입이 서로 다른 변수끼리의 연산시 형변화 과정에서 오차가 발생할 수 있다. -> 치명적인 버그가 될 수 있음
	// Golang 에서는 자료형이 다른 변수끼리의 연산을 아예 금지한다 (예외 발생)
	// + - * % / << >> & ^

	// 일반 integer 는 uinteger 의 약 절반 정도의 범위를 갖는다.
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println("n1 : ", n1)
	fmt.Println("n2 : ", n2)
	fmt.Println("n3 : ", n3)
	fmt.Println("n4 : ", n4)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	n5 := 100000
	n6 := int16(10000)
	n7 := uint8(100)

	// 자료형이 다른 변수 간에 연산을 할 때에는(논리연산자 포함) 반드시 형 변환을 통해 같은 타입으로 맞춰줘야 한다.
	fmt.Println(n5 + int(n6))
	fmt.Println(n6 + int16(n7))
	fmt.Println(n6 > int16(n7))
}
