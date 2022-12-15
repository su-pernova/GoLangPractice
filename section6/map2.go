package main

import "fmt"

func main() {
	// 맵의 원소를 조회하고 순회할 수 있다.
	map1 := map[string]string{
		"daum":   "https://daum.net",
		"naver":  "https://naver.com",
		"google": "https://google.com",
	}

	fmt.Println(map1["daum"])
	fmt.Println(map1["naver"])
	fmt.Println(map1["google"], "\n")

	// 일반적인 배열, 슬라이스와 달리(index, value 쌍 반환) range map 은 key, value 쌍을 반환한다.
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println()

	for _, v := range map1 {
		fmt.Println(v)
	}
}
