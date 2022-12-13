package main

import (
	"fmt"
	"math"
)

func main() {
	// 언더플로우, 오버플로우가 발생하는 경우를 확인할 수 있다.
	//var n1 uint8 = math.MaxUint8 + 1
	//var n2 uint16 = math.MaxUint16 + 1
	//var n3 uint32 = math.MaxUint32 + 1
	var n4 uint64 = math.MaxUint64 + 1 // 이것도 에러 나는데 IDE 에서는 안잡아주네용
	fmt.Println(n4)
}
