package main

import "fmt"

func main() {
	// 변수 종류 : 정수, 실수(소수점), 문자열, Boolean
	// 변수명 : 숫자 첫글자 불가, 대소문자 구분, 문자, 숫자, 밑줄, 특수기호 사용 가능
	// 변수 및 상수 위치 : 함수 안,밖에서 사용 가능

	// 변수를 선언할 수 있다.
	var a int
	var b string
	var c, d, e int

	// 변수에 값을 할당할 수 있다.
	a = 4
	b = "Hi! Go!"
	c, d, e = 1, 2, 3

	// 선언과 초기화를 동시에 할 수 있다.
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!"

	// 자료형을 지정하지 않아도 값을 할당할 수 있다.
	var k = 4.75
	var l = "Hi! Go Language!"
	var m = true

	// Go 에서는 변수를 선언하고 사용하지 않으면 에러가 난다.
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c, d, e : ", c, d, e)
	fmt.Println("f, g, h : ", f, g, h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)
}
