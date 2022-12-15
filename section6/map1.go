package main

import "fmt"

func main() {
	// 맵(해시테이블, 딕셔너리 이라 부르기도 한다.)
	// Key-Value 로 자료를 저장한다.
	// 참조값을 전달하는 레퍼런스 타입이다. -> 비교연산자 사용이 불가능하다.
	// Key 값에는 참조타입 변수를 사용할 수 없다. Value 값에는 변수 타입 상관없이 모두 사용할 수 있다.
	// make 함수 및 축약(리터럴)로 초기화할 수 있다.
	// (주로 키 값으로 조회하기 때문에) 원소의 순서가 정해져있지 않다.

	// 1) 가장 정석적인 선언 방법
	// 대괄호 안에 key 의 자료형을, 그 뒤에 value 의 자료형을 적는다.
	var map1 map[string]int = make(map[string]int)

	// 2) 조금 더 간단한 선언 방법
	var map2 = make(map[string]int)

	// 3) 짧은 선언. 가장 일반적으로 사용된다.
	map3 := make(map[string]int)
	map4 := map[string]int{}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println(map4, "\n")

	// map 선언 마지막에 중괄호 {} 를 붙여 JSON 형태로 감쌀 수 있다.
	// 선언 후 값을 추가하는 방식
	map5 := map[string]int{}
	map5["apple"] = 10
	map5["banana"] = 20
	map5["orange"] = 30

	map6 := map[string]int{
		"apple":  10,
		"banana": 40,
		"orange": 23,
	}

	map7 := make(map[string]int, 10)
	map7["apple"] = 10
	map7["banana"] = 20
	map7["orange"] = 30

	fmt.Println(map5)
	fmt.Println(map6)
	fmt.Println(map7, "\n")

	fmt.Println("map7[\"apple\"] : ", map7["apple"])
	fmt.Println("map7[\"orange\"] : ", map7["orange"])
}
