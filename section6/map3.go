package main

import "fmt"

func main() {
	// 맵에 원소를 추가하거나, 기존 원소를 수정/삭제할 수 있다.
	map1 := map[string]string{
		"daum":   "https://daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
		"test1":  "https://test1.com",
	}

	fmt.Println(map1)

	// 맵 원소 추가
	map1["test2"] = "https://test2.com"
	fmt.Println(map1)

	// 맵 원소 수정
	map1["test2"] = "https://test02.com"
	fmt.Println(map1)

	// 맵 원소 삭제
	delete(map1, "test2")
	fmt.Println(map1)
}
