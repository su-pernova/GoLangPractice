package main

import "fmt"

func main() {
	// switch-case 문의 case 표현식에서 여러개의 변수를 사용할 수 있다.
	// fallthrough 예약어를 사용할 수 있다.

	a := 30 / 15
	switch a {

	// a 가 2, 4, 6 인 경우
	case 2, 4, 6:
		fmt.Println("a -> ", a, " : 짝수이다.")

	// a 가 1, 3, 5 인 경우
	case 1, 3, 5:
		fmt.Println("a -> ", a, " : 홀수이다.")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("e 는 java 이다.")
		// 해당 case 문이 조건에 맞아 실행되었다면 바로 그 다음 case 문도 조건이 맞는 것으로 간주된다.
		fallthrough
	case "go":
		fmt.Println("e 는 go 이다.")
		fallthrough
	case "python":
		fmt.Println("e 는 python 이다.")
		fallthrough
	case "ruby":
		fmt.Println("e 는 ruby 이다.")
		fallthrough
	case "c":
		fmt.Println("e 는 c 이다.")
	}

}
