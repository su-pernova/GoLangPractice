package main

import "fmt"

func main() {
	// Golang 문자열 표현의 특징을 이해할 수 있다.
	// UTF-8 인코딩 : 한글, 영어, 일본어 크게 제약 없이 사용할 수 있다.
	// 사용시에는 바이트 수를 주의하자.

	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	// 문자열의 각 문자는 메모리에 순차적으로 저장된다.
	// 따라서 하나의 문자열은 마치 하나의 배열처럼 사용할 수 있다. (인덱스 등)
	// 아래처럼 출력하면 문자열의 아스키코드가 찍힌다.
	fmt.Println(str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println(str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println(str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// 아래처럼 해주어야 아스키코드가 아닌 실제 문자가 출력된다.
	fmt.Printf("%c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("%c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("%c %c %c %c %c %c\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])

	// 한글의 경우에는 아래와 같이 처리해 주어야 문자가 깨지지 않는다.
	conStr := []rune(str3)
	fmt.Printf("%c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	// 아래와 방식으로 문자열을 순회하는 것은 매우 자주 사용되는 패턴이므로 잘 기억해두자.
	for i, char := range str1 {
		fmt.Printf("%c(%d)\t", char, i)
	}
	fmt.Println()

	for i, char := range str2 {
		fmt.Printf("%c(%d)\t", char, i)
	}

}
