package main

import "fmt"

func main() {
	// 제어문(조건문) : if - else if - else 문을 사용할 수 있다.
	i := 40

	if i >= 120 {
		fmt.Println("i 는 120 이상이다.")
	} else if i < 120 && i >= 100 {
		fmt.Println("i 는 120 미만 100 이상이다.")
	} else if i < 100 && i >= 50 {
		fmt.Println("i 는 100 미만 50 이상이다.")
	} else {
		fmt.Println("i 는 50 미만이다.")
	}
}
