package main

import "fmt"

func main() {
	// 배열의 복사 : 참조 복사가 아닌 값 복사가 일어남을 이해할 수 있다.

	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Printf("%v / %p\n", arr1, &arr1)
	fmt.Printf("%v / %p\n", arr2, &arr2)
}
