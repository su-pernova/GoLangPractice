package main

import "fmt"

func main() {
	// 문자열 비교 연산 예제
	str1 := "Golang" // 6 byte
	str2 := "World"  // 5 byte

	fmt.Println(str1 == str2)
	fmt.Println(str1 != str2)

	// Golang(6B)이 World(5B) 보다 클 것 같지만, 비교 연산 결과는 World 가 더 크다
	// Go 문자열의 비교 연산은 아스키 코드에 대한 사전식 비교이기 때문
	fmt.Println(str1 > str2)
	fmt.Println(str1 < str2)

	// 대문자 vs 소문자 : 소문자의 아스키 코드 값이 더 크기 때문에 대소비교시 소문자가 더 크다.
	str3 := "A"
	str4 := "a"
	fmt.Println("A , a : ", str3[0], str4[0])
	fmt.Println("A < a : ", str3[0] < str4[0])

}
