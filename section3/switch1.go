package main

import "fmt"

func main() {
	// 제어문(조건문) : switch - case 문을 작성할 수 있다.
	// switch 뒤 표현식(Expression) 생략 가능
	// case 뒤 표현식(Expression) 사용 가능
	// 자동 break 때문에 fallthrough 존재
	// 값이 아닌 변수 자료형으로 분기 처리 가능

	// 함수 영역 안에서 사용하는 변수를 함수 영역 밖에서 선언했다 : 일반적이지 않다.
	a := -7
	switch {
	case a < 0:
		fmt.Println("a 는 음수이다.")
	case a == 0:
		fmt.Println("a 는 0 이다.")
	case a > 0:
		fmt.Println("a 는 양수이다.")
	}

	// 함수 영역 안에서 사용하는 변수는 짧은 선언으로 하는 것이 일반적이다.
	switch b := 27; {
	case b < 0:
		fmt.Println("b 는 음수이다.")
	case b == 0:
		fmt.Println("b 는 0 이다.")
	case b > 0:
		fmt.Println("b 는 양수이다.")
	}

	// switch 뒤 표현식에 내용을 추가해서 아래와 같이 사용할 수 있다.
	switch c := "go"; c {
	case "go":
		fmt.Println("c 는 go 이다.")
	case "java":
		fmt.Println("c 는 java 이다.")
	default:
		fmt.Println("c 와 일치하는 값이 없다.")
	}

	// switch 뒤 표현식에서 연산을 수행할 수 있다.
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("c 는 golang 이다.")
	case "javalang":
		fmt.Println("c 는 javalang 이다.")
	default:
		fmt.Println("c 와 일치하는 값이 없다.")
	}

	// 짧은 선언 뒤에는 세미콜론이 와야 한다는 점을 유의하자.
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i 는 j 보다 작다.")
	case i == j:
		fmt.Println("i 와 j 는 같다.")
	case i > j:
		fmt.Println("i 는 j 보다 크다.")
	}

}
