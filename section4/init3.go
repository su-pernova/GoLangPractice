package main

import (
	"fmt"
	// lib 패키지를 호출하면서 해당 패키지의 init 함수가 가장 먼저 실행된다.
	"goLangPractice/section4/lib"
)

var num int32

// 그 다음으로 현재 패키지의 init 함수가 실행된다.
// 목적 : 변수 초기화
func init() {
	num = 30
	fmt.Println("Main init start!")
}

// 마지막으로 현재 패키지의 main 함수가 실행된다.
func main() {
	fmt.Println("10 보다 큰 수인가?", lib.CheckNum(num))
}
