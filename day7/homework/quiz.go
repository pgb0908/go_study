package homework

import "strings"

// Q1. n번째 피보나치 수를 반환합니다.
// F(0)=0, F(1)=1, F(n)=F(n-1)+F(n-2)
func Fibonacci(n int) int {
	// TODO
	return 0
}

// Q2. 문자열을 순서대로 처리합니다:
// 1. 앞뒤 공백 제거  2. 소문자 변환  3. 공백을 '_'로 교체
func ProcessString(s string) string {
	// TODO: strings 패키지 활용
	_ = strings.TrimSpace // 이 줄은 삭제하세요
	return ""
}

// Q3. 최고 점수 학생 이름과 전체 평균을 반환합니다.
// 빈 맵이면 ("", 0.0) 반환
func GradeReport(scores map[string]int) (string, float64) {
	// TODO
	return "", 0.0
}
