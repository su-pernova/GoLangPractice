package main

import (
	"fmt"
	"goLangPractice/section4/lib2"
)

func main() {
	// 패키지 접근 제어를 할 수 있다.
	// 변수, 상수, 함수, 메서드, 구조체 등의 식별자에 접근할 수 있다.
	// 첫 글자가 대문자 : 패키지 외부에서 접근 가능 (public)
	// 첫 글자가 소문자 : 패키지 외부 접근 불가 (패키지 내에서만 접근 가능)(protected)

	fmt.Println("100 보다 큰 수인가? : ", lib2.CheckNum1(101))
	fmt.Println("1000 보다 큰 수인가? : ", lib2.CheckNum2(999))
}
