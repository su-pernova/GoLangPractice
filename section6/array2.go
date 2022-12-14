package main

import "fmt"

func main() {
	// 배열의 값을 순회할 수 있다.

	arr1 := [5]int{1, 10, 100, 1000, 10000}
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	fmt.Println()

	// 조금 더 일반적으로 사용하는 방식이다.
	arr2 := [5]int{7, 77, 777, 7777, 77777}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

}
