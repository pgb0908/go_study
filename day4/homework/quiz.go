package homework

import "strings"

// Q1. 중복을 제거하고 처음 등장 순서를 유지해 반환합니다.
func Unique(nums []int) []int {
	// TODO: 맵을 활용하세요
	return nil
}

// Q2. 문자열을 공백으로 분리해 단어 빈도를 맵으로 반환합니다.
func WordCount(s string) map[string]int {
	// TODO: strings.Fields(s) 로 단어 분리 가능
	_ = strings.Fields // 이 줄은 삭제하세요
	return nil
}

// Q3. Rectangle 구조체를 정의하세요. (TODO: Width, Height float64 필드 추가)
type Rectangle struct {
	Width  float64 // TODO: 그대로 두거나 직접 작성해보세요
	Height float64
}

// NewRectangle: Rectangle 생성자
func NewRectangle(w, h float64) Rectangle {
	// TODO
	return Rectangle{}
}

// Area: 넓이 반환
func Area(r Rectangle) float64 {
	// TODO
	return 0
}

// Perimeter: 둘레 반환
func Perimeter(r Rectangle) float64 {
	// TODO
	return 0
}

// Q4. 포인터로 전달된 값을 2배로 만듭니다.
func Double(n *int) {
	// TODO
}
