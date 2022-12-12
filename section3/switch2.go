package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Go 에서는 다른 언어와 달리 switch - case 문을 조금 더 다양한 형태로 사용할 수 있다.
	// if 문에서 처럼 && 연산자 등을 사용할 수 있다.

	rand.Seed(time.Now().UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i -> ", i, " : 50 이상 100 미만")
	case i >= 25 && i < 50:
		fmt.Println("i -> ", i, " : 25 이상 50 미만")
	default:
		fmt.Println("i -> ", i, " : 아무 조건에도 해당되지 않는다.")
	}

}
