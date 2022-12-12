package main

import "fmt"

func main() {
	// 반복문 for 를 활용할 수 있다.

	// 1 부터 100 까지의 합을 계산하는 for 문
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += i
	}
	fmt.Println("sum1 : ", sum1)

	// 1 부터 100 까지의 합을 계산하는 for 문의 다른 형태
	// 주로 이 형태를 조금 더 많이 사용한다.
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		// j = i++ // 이 형태로 사용할 수 없다. Go 에서 후치연산은 반환값이 없기 때문
	}
	fmt.Println("sum2 : ", sum2)

	// 무한루프 형태로 사용할 때 break 를 사용할 수 있다.
	sum3, i := 0, 0
	for {
		if i > 100 {
			break // 가장 가까운 반복문을 탈출한다.
		}
		sum3 += i
		i++
	}
	fmt.Println("sum3 : ", sum3)

	// 여러개의 변수에 변화를 가할 수 있다.
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("i, j : ", i, j)
	}

	/*
		[ 에러가 발생하는 경우 ]
		1. 맨 마지막 칸 처럼 작성하면 에러가 발생한다. 즉 바로 위의 예제처럼 사용해야 한다.
			for i, j := 0, 0; i <= 10; i++, j+= 10 {
				fmt.Println("i, j : ", i, j)
			}
	*/

}
