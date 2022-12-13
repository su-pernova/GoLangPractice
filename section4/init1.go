package main

import "fmt"

// 패키지 로드 시에 가장 먼저 (main 함수보다 먼저) 딱 한 번 실행되는 함수이다.
// 변수 설정 또는 프로그램 상태 체크 등 가장 먼저 초기화가 필요한 작업에 사용할 수 있다.
func init() {
	fmt.Println("Init Method Start!")
}

// 만약 main 메소드가 없는 파일에 init 메소드가 있다면, 해당 초기화 메소드는 해당 파일이 호출되는 시점에 실행된다.
func main() {
	fmt.Println("Main Method Start!")
}
