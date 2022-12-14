package main

import (
	"fmt"
	"sort"
)

func main() {
	// 슬라이스를 추출 및 정렬할 수 있다.
	// < 추출 >
	//	- slice[i:j] = i 번째 인덱스 부터 j-1 번째 인덱스까지 추출
	//	- slice[i:] = i 번째 인덱스부터 마지막까지 추출
	//	- slice[:j] = 처음부터 j-1 번째 인덱스까지 추출
	//	- slice[:] =  처음부터 마지막까지 추출

	// 슬라이스 추출 예제
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice1[:])
	fmt.Println(slice1[0:])
	fmt.Println(slice1[:3])
	fmt.Println(slice1[3:])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[1:3])
	fmt.Println(slice1[0:len(slice1)], "\n")

	// 슬라이스 정렬 예제
	// sort 패키지 참고문서 : https://golang.org/pkg/sort
	slice2 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}

	fmt.Print(slice2)
	fmt.Println(" / 정렬 유무 : ", sort.IntsAreSorted(slice2)) // 정렬이 되어있는지 여부를 확인(T/F)
	sort.Ints(slice2)
	fmt.Print(slice2)
	fmt.Println(" / 정렬 유무 : ", sort.IntsAreSorted(slice2), "\n")

	slice3 := []string{"b", "d", "f", "a", "c", "e"}

	fmt.Print(slice3)
	fmt.Println(" / 정렬 유무 : ", sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Print(slice3)
	fmt.Println(" / 정렬 유무 : ", sort.StringsAreSorted(slice3))

}
