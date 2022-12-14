package main

import "fmt"

func main() {
	// 슬라이스의 복사 : 값 복사가 아닌 참조 복사가 일어남을 이해할 수 있다.
	// 슬라이스를 순회할 수 있다.

	// 배열의 복사는 값 복사이기 때문에 복사 후 한쪽의 값을 변경하면 해당 변경 사항이 그 쪽에만 반영된다.
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1

	arr2[0] = 7

	fmt.Println(arr1)
	fmt.Println(arr2)

	// 슬라이스의 복사는 참조 복사이기 때문에 복사 후 한쪽의 값을 변경하면 해당 변경 사항이 두 쪽 모두에 반영된다.
	slice1 := []int{1, 2, 3}
	slice2 := slice1

	slice2[0] = 7

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]int, 50, 100)
	fmt.Println(slice3[4])

	// 슬라이스 순회 예제
	for i, v := range slice1 {
		fmt.Println(i, v)
	}

}
