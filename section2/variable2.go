package main

import "fmt"

func main() {
	// 변수를 여러개 선언하는 패턴
	// 관련 있는 변수끼리 묶어 선언함으로서 가독성을 높일 수 있다.
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	// 묶어서 선언했더라도 밖에서 초기화할 수 있다.
	height = 250
	weight = 350.56
	isRunning = true

	fmt.Println("name : ", name, " / height : ", height, " / weight : ", weight, " / isRunning : ", isRunning)
}
