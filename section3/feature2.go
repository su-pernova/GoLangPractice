package main

import "fmt"

func main() {
	// 문장 끝 세미콜론이 자동으로 생성됨을 항상 주의하자.
	// 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론을 사용해야 하나, 권장하지 않는다.
	// 반복문 및 제어문(if, for)을 사용할 때 특히 주의하자.

	for i := 0; i <= 10; i++ {
		fmt.Print("i : ")
		fmt.Println(i)
	}

	for j := 10; j >= 0; j-- {
		fmt.Println("j : ", j)
	}
}
