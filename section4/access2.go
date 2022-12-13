package main

import (
	"fmt"
	checkUp "goLangPractice/section4/lib"
	_ "goLangPractice/section4/lib2"
)

func main() {
	// 패키지 접근제어
	// 별칭 사용 (파이썬의 import as 와 동일)
	// 빈 식별자(_) 사용하기 : 사용하지 않는 package 를 그냥 두면 컴파일 시 사라지지만, 빈 식별자 처리를 해두면 사라지는 것을 방지할 수 있다.

	fmt.Println("10 보다 큰 수인가? : ", checkUp.CheckNum(12))
}
