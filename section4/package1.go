// Go 에서 main 패키지는 특별하게 인식된다.
// 컴파일러가 main 패키지로 선언된 go 파일은 공유 라이브러리가 아닌, 프로그램의 시작점으로 인식한다.
// 즉 main 함수가 사용되는 파일의 package 는 main 이어야 한다.
package main

import (
	"fmt"
	"os"
)

func main() {
	// 패키지를 활용해 코드 구조화(모듈화) 및 재사용을 할 수 있다.
	// 	- 응집도는 높이고, 결합도는 낮추는 것이 좋다.
	// 	- 패키지 단위의 독립적이고 작은 단위로 개발할 것
	// 	- 작은 패키지를 결합해서 프로그램을 작성할 것
	// 구조 설정 규칙
	// 	- 패키지 이름 = 디렉토리 이름
	// 	- 같은 패키지 내 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
	// 함수 네이밍 규칙
	// 	- 소문자 private
	// 	- 대문자 public

	var name string
	fmt.Println("이름을 입력하세요 : ")
	fmt.Scanf("%s", &name)
	fmt.Fprintf(os.Stdout, "%s 님 반갑습니다.\n", name)

}
