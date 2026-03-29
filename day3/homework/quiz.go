package homework

// Q1. 슬라이스에서 최솟값과 최댓값을 반환합니다.
// 반환: (min, max)
func MinMax(nums []int) (int, int) {
	// TODO
	return 0, 0
}

// Q2. 가변 인자를 모두 더해 반환합니다.
func Sum(nums ...int) int {
	// TODO
	return 0
}

// Q3. 호출마다 1씩 증가하는 카운터 함수를 반환합니다.
func MakeCounter() func() int {
	// TODO: 클로저 사용
	return nil
}

// Q4. defer 등록 순서: "first", "second", "third"
// 실제 실행 순서(LIFO)대로 슬라이스에 담아 반환합니다.
func DeferOrder() []string {
	result := []string{}
	// TODO: defer를 3번 사용
	// 힌트: defer func() { result = append(result, "...") }()
	return result
}
