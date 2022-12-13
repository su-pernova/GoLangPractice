package main

import "fmt"

// init 메소드는 여러개 존재할 수 있고, 위에서부터 차례대로 실행된다.
// 하지만 초기화 메소드가 여러개 있는건 바람직한 형태의 코드는 아니다.
func init() {
	fmt.Println("Init Method 1")
}

func init() {
	fmt.Println("Init Method 2")
}

func init() {
	fmt.Println("Init Method 3")
}

func init() {
	fmt.Println("Init Method 4")
}

func main() {
	fmt.Println("Main Method Start!")
}
