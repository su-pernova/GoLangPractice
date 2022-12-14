package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열 데이터 타입을 사용할 수 있다.
	// 큰 따옴표(""), 백스쿼트(``) 를 사용해 문자열을 표현할 수 있다.
	// 작은 따옴표('') 를 사용해 문자를 표현할 수 잇다.
	// Golang 에는 char 타입이 존재하지 않는다.
	// 대신 문자 코드 값과 rune(int32) 자료형으로 문자를 표현할 수 있다.
	// 자주 사용하는 escape : \\, \', \" (따옴표 안에서 역슬래쉬 뒤에 오는 문자는 문자 그대로 찍힌다.)
	// \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭) 등

	// str1 과 str2 의 문자열은 동일하다.
	var str1 string = "c:\\go_study\\src\\" // c:\go_study\src\
	str2 := `c:\go_study\src\`              // 백스쿼트 내부의 문자열은 작성된 내용 그대로 채택된다. 백슬래시를 사용할 때 보다 더 나은 가독성을 유지할 수 있다.
	fmt.Println(str1)
	fmt.Println(str2)

	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요."
	var str5 string = "\ud55c\uae00"
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)

	// 문자열의 크기(byte)를 확인할 수 있다. (물리적 길이 아님!)
	fmt.Println("len(str3) : ", len(str3))
	fmt.Println("len(str4) : ", len(str4))

	// 물리적 길이는 ~를 활용하여 구할 수 있다.
	fmt.Println("str3 의 길이 : ", utf8.RuneCountInString(str3))
	fmt.Println("str4 의 길이 : ", utf8.RuneCountInString(str4))
	fmt.Println("str4 의 길이 : ", len([]rune(str4))) // len 을 사용해서 길제 길이를
}
