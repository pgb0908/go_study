package homework

import "fmt"

// Q1. 아래 함수를 완성하세요.
// var 키워드, := 단축선언, var 블록 3가지를 모두 사용해야 합니다.
func DeclareVariables() (string, int, float64, float64) {
	// TODO: var 키워드로 name 선언 (값: "Gopher")

	// TODO: := 로 age 선언 (값: 5)

	// TODO: var 블록으로 height(171.5), weight(68.0) 선언

	return "", 0, 0, 0 // TODO: 선언한 변수로 교체하세요
}

// Q2. Weekday 타입과 iota 상수를 선언하세요.
type Weekday int

const (
// TODO: Mon=1 부터 Sun=7 까지 iota로 선언
)

// IsWeekend: Sat 또는 Sun이면 true를 반환합니다.
func IsWeekend(d Weekday) bool {
	// TODO
	return false
}

// Q3. 섭씨를 화씨로 변환합니다. 공식: F = C*9/5 + 32
func CelsiusToFahrenheit(c float64) float64 {
	// TODO
	return 0
}

// Q4. 프로필 문자열을 반환합니다.
// 형식: "이름: Gopher, 나이: 5, 점수: 98.50"
func FormatProfile(name string, age int, score float64) string {
	// TODO: fmt.Sprintf 사용
	_ = fmt.Sprintf // 이 줄은 삭제하세요
	return ""
}
