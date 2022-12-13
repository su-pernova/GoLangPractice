package lib2

// 이렇게 main 메소드가 없는 파일은 단독으로 실행될 수 없다.

func CheckNum1(c int32) bool {
	return c > 100
}

func CheckNum2(c int32) bool {
	return c > 1000
}
