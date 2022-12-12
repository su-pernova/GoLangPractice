package main

import "fmt"

func main() {
	const a, b int = 10, 99
	// 하나의 줄에서 서로 다른 자료형의 상수를 여러개 선언, 초기화할 수 있다.
	const c, d, e = true, 0.84, "test"
	const (
		x, y int16 = 50, 90
		i, k       = "Data", 7776
	)

	fmt.Println("a, b, c, d, e : ", a, b, c, d, e)
	fmt.Println("x, y, i, k : ", x, y, i, k)
}
