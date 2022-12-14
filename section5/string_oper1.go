package main

import "fmt"

func main() {
	// 문자열 자료형의 연산을 수행할 수 있다.
	// 문자열 추출, 비교, 결합
	// Golang 의 문자열 연산은 파이썬만큼 간편한다!

	// 문자열 추출 연산 예제
	var str1 string = "Golang"
	var str2 string = "World"

	// 하나의 문자만 출력하면 아스키 코드로, 구간을 출력하면 string 값 그대로 출력된다.
	// string 은 있지만 char 는 없어서? 그런것? 같다?
	fmt.Println(str1[0:2], str1[0])
	fmt.Println(str2[3:], str2[0])
	fmt.Println(str2[:4])
	fmt.Println(str1[1:3])

	// 전자는 W의 아스키 코드가, 후자는 W 값이 그대로 출력된다.
	fmt.Println(str2[0], str2[0:1])
}
