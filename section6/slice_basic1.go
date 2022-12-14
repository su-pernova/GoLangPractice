package main

import "fmt"

func main() {
	// Go 에서 슬라이스를 활용할 수 있다.
	// 슬라이스는 길이가 고정적이지 않고, 동적으로 크기가 늘어날 수 있다.
	// 래퍼런스(참조) 타입이다.
	// 슬라이스는 길이와 용량의 개념이 있고, 크기를 동적으로 할당할 수 있다.
	// switch 를 선언하기 위한 방법으로는 2가지 정도가 있다.
	//	- 배열처럼 선언하는 방법
	// 	- make 함수를 활용하는 방법 : make(자료형, 길이, 용량(생략시 길이 값과 동일한 값으로 설정됨))

	// 배열과 비슷하게 선언하는 방법
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	slice3[4] = 10 // 값을 수정할 수 있다.

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	// make 함수를 이용해 선언하는 방법(make(자료형, 길이, 용량))
	// < 용량 설정은 왜 하는 것일까? >
	//	- 아래 예시처럼 하면 10의 용량을 할당받아 두고, 5의 용량만 먼저 사용하게 된다.
	//		- 10의 용량을 미리 할당받아 두지 않았다면, 5의 용량을 다 사용한 이후 원소를 추가했을 때 새로운 공간을 재할당 받게 된다. (재할당 과정에서 지연이 발생할 수 있다)
	//		- 10의 용량을 미리 할당받아 두었다면, 5의 용량을 다 사용한 이후 원소를 추가해도 새로운 공간을 할당받지 않는다.
	//		- 너무 많이 받아두고 조금만 쓰는 것도 메모리 낭비이기 때문에 좋지 않다.
	//	- 즉 적당한 크기의 용량(너무 많지도, 적지도 않게)을 할당받는 것이 중요하다.
	// < 유의사항 >
	//	- 0의 값으로 자동 초기화는 선언한 길이 만큼만 수행된다.
	//	- 즉 길이가 10이고 용량이 15일 경우, 0이 10개 저장되므로 index 13 을 찍으면 index out of bound 에러가 발생한다.
	var slice5 []int = make([]int, 5, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7

	fmt.Printf("\n%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	var slice9 []int // 길이와 용량이 0인 슬라이스 : nil 이라고 한다.
	if slice9 == nil {
		fmt.Println("\nThis is Nil Slice")
	}

}
