package main

import "fmt"

// 가변 인자 함수: ...타입 형식으로 선언
// 인자 개수를 미리 정하지 않아도 됩니다.
// nums는 함수 내부에서 []int 슬라이스처럼 동작합니다.
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 가변 인자와 일반 인자를 함께 사용
func greetAll(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

func main() {
	fmt.Println("=== 가변 인자 함수 ===")

	// 인자 개수를 자유롭게 전달
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4, 5))

	// 인자를 아예 안 넣어도 됩니다
	fmt.Println(sum()) // 0

	// =========================================================
	// 슬라이스를 펼쳐서 가변 인자로 전달
	// 슬라이스 변수 뒤에 ... 을 붙입니다
	// =========================================================
	fmt.Println("\n=== 슬라이스 펼치기 ===")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("슬라이스: %v\n", numbers)
	fmt.Printf("합계: %d\n", sum(numbers...)) // numbers... 로 펼쳐서 전달

	// =========================================================
	// fmt.Println 자체가 가변 인자 함수입니다!
	// func Println(a ...any) (n int, err error)
	// =========================================================
	fmt.Println("\n=== fmt.Println은 가변 인자! ===")
	fmt.Println("하나")
	fmt.Println("하나", "둘", "셋")
	fmt.Println(1, 2.5, true, "문자열")

	// 일반 인자 + 가변 인자
	fmt.Println("\n=== 혼합 인자 ===")
	greetAll("안녕하세요", "철수", "영희", "민수")
}
