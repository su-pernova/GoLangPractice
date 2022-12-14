package main

import (
	"fmt"
	"strings"
)

func main() {
	// 문자열 결합 연산 예제
	// 일반연산(+) 또는 Join 연산을 사용할 수 있다.
	// Join 연산을 사용하는 것이 더 효율이 좋다. (시간/공간복잡도)
	// 따라서 실 서비스 개발시에는 반드시 Join 을 사용하도록 하자.

	// 일반연산(+)을 사용하는 경우
	str1 := "Hello " +
		"World " +
		"Nice to Meet You "
	str2 := "Good Morning"

	fmt.Println(str1 + str2)

	// append, Join 을 사용하는 경우
	strSet := []string{} // 중괄호 : string 슬라이스 자료형
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println(strings.Join(strSet, "----- ")) // 두번째 매개변수 = strSet 에 들어있는 2개 이상의 문자열 결합시 삽입되는 구분자
}
