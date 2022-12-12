package main

import "fmt"

func main() {
	// 반복문 for 를 사용할 수 있다.
	// Go 는 다른 언어와 달리, 반복문이 for 키워드만 유일하게 제공된다.
	// for 문을 사용해서 무한루프를 구현할 수 있다.

	for i := 0; i < 5; i++ {
		fmt.Println("i 값은 ", i, "이다.")
	}

	/*
		[ 에러가 발생하는 경우 ]
		1. if, else, switch 에서와 동일하게 중괄호 이전에 줄바꿈이 들어가면 에러가 발생한다.
			for i := 0; i < 5; i++
			{
				fmt.Println("i 값은 ", i, "이다.")
			}

		2. 동일하게 중괄호를 생략해도 에러가 발생한다.
			for i := 0; i < 5; i++
				fmt.Println("i 값은 ", i, "이다.")
	*/

	// for 문에 조건을 두지 않으면 무한 루프가 된다.
	//for {
	//	fmt.Println("Infinite Loop")
	//}

	// 배열에서 index 와 value를 가져올 수 있다.
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println(index, name)
	}
	for _, name := range loc { // 이런식으로 묵음 처리도 가능하다.
		fmt.Println(name)
	}
}
