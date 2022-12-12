package main

import "fmt"

func main() {
	// 언더바(_)는 사용할 수 없는 변수이다.
	// 즉 iota 사용시 건너뛰고 싶은 부분은 _ 를 사용하면 된다.
	const (
		_ = iota
		A
		B
		C
	)

	fmt.Println("A, B, C : ", A, B, C)
	// fmt.Println("_ : ", _) // 에러 발생 : 언더바(_)는 사용할 수 없는 변수이기 때문
}
