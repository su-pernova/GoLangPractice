package main

import "fmt"

func main() {
	// 맵의 원소를 조회할 때의 유의사항을 이해할 수 있다.

	map1 := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 1115,
		"lemon":  0,
	}

	// 존재하지 않는 Key를 조회하는 경우, 기본 자료형의 초기화값이 반환된다.
	// 즉 value 값만 조회하는 경우 이게 key 가 존재하지 않는 것인지, 실제 저장된 value 값이 0 인지를 구분할 수 없다.
	keyvalue1 := map1["lemon"]
	keyvalue2 := map1["kiwi"]

	fmt.Println(keyvalue1, keyvalue2)

	// 위 상황을 구분하기 위해, Key 값에 대한 조회를 두개의 인스턴스로 받아 확인할 수 있다.
	// 첫번째 인스턴스에는 value 값이, 두번째 인스턴스에는 key 값의 존재 여부(bool)가 반환된다.
	value1, exist1 := map1["kiwi"]
	value2, exist2 := map1["apple"]

	fmt.Println(value1, exist1, "/", value2, exist2, "\n")

	// 이는 조건문에서 유용하게 활용할 수 있다.
	// 아래와 같은 상황에서 유용 (코테에서 자주 마주했던...)
	//	- 맵에 특정 키가 있다면 해당 키의 value 에 값을 저장하라.
	//	- 없다면 키를 생성하고, 해당키의 value 에 값을 저장하라.
	fmt.Println("before :", map1)
	if value, exist := map1["kiwi"]; exist {
		fmt.Println(value)
	} else {
		map1["kiwi"] = 50
	}
	fmt.Println("after :", map1)

}
