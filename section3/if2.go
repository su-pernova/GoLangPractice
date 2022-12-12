package main

import "fmt"

func main() {
	// 제어문(조건문) : if - else 문을 사용할 수 있다.
	var a int = 50
	b := 70

	if a >= 65 {
		fmt.Println("a 는 65 이상이다.")
	} else {
		fmt.Println("a 는 65 미만이다.")
	}

	if b >= 70 {
		fmt.Println("b 는 70 이상이다.")
	} else {
		fmt.Println("b 는 70 미만이다.")
	}

	/*
		[ 에러가 발생하는 경우 ]
		1. if 문 에서와 같이 else 문 중괄호 이전에 줄바꿈이 들어가면 에러가 발생한다.
			if a >= 65 {
				fmt.Println("a 는 65 이상이다.")
			} else
			{
				fmt.Println("a 는 65 미만이다.")
			}

		2. else 문 선언 이전에 줄바꿈이 들어가도 에러가 발생한다.
			if a >= 65 {
				fmt.Println("a 는 65 이상이다.")
			}
			else {
				fmt.Println("a 는 65 미만이다.")
			}
	*/
}
