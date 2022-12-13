package main

import (
	"fmt"
	"goLangPractice/section4/lib"
)

func main() {
	// 직접 작성한 패키지를 import 하여 사용할 수 있다.
	fmt.Println("10 보다 큰 수인가? : ", lib.CheckNum(14))
}
