package main

import "fmt"

func main() {
	// break 문을 사용할 때 추가 키워드를 사용하면 가장 가까운 반복문이 아닌, 더 먼 반복문을 탈출할 수 있다.
	// 반복문에서 continue 키워드를 사용할 수 있다.
	// continue 문을 사용할 때 추가 키워드를 사용하면 가장 가까운 반복문이 아닌, 더 먼 반복문으로 이동할 수 있다.

	// 두번째로 가까운 반복문을 탈출한다.
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println(i, j)
		}
	}

	// 홀수만 출력하는 반복문
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i : ", i)
	}

	// 두번째로 가까운 반복문으로 이동한다.
Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("i, j : ", i, j)
		}
	}
}
