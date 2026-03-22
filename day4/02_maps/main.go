package main

import "fmt"

func main() {
	// =========================================================
	// 맵 (Map): key-value 쌍을 저장하는 자료구조
	// map[키타입]값타입
	// =========================================================
	fmt.Println("=== 맵 생성 ===")

	// 방법 1: 리터럴로 초기화
	prices := map[string]int{
		"사과":  1500,
		"바나나": 800,
		"딸기":  3000,
	}
	fmt.Println("과일 가격:", prices)

	// 방법 2: make로 생성 (나중에 채울 때)
	scores := make(map[string]int)
	scores["철수"] = 90
	scores["영희"] = 85
	scores["민수"] = 78
	fmt.Println("학생 점수:", scores)

	// =========================================================
	// 맵 읽기, 수정, 삭제
	// =========================================================
	fmt.Println("\n=== CRUD ===")
	fmt.Printf("사과 가격: %d원\n", prices["사과"])

	prices["사과"] = 2000 // 수정
	fmt.Printf("사과 가격 수정 후: %d원\n", prices["사과"])

	delete(prices, "바나나") // 삭제
	fmt.Println("바나나 삭제 후:", prices)

	// =========================================================
	// 존재 여부 확인: ok 패턴 (매우 중요!)
	// 없는 키를 접근하면 에러가 아닌 제로값이 반환됩니다.
	// ok로 실제 존재하는지 확인해야 합니다.
	// =========================================================
	fmt.Println("\n=== ok 패턴 ===")
	price, ok := prices["바나나"] // 삭제했으니 없음
	if ok {
		fmt.Printf("바나나: %d원\n", price)
	} else {
		fmt.Println("바나나 가격 정보 없음 (ok=false, price=", price, ")")
	}

	price, ok = prices["사과"]
	if ok {
		fmt.Printf("사과: %d원\n", price)
	}

	// =========================================================
	// range로 맵 순회 (순서는 보장되지 않습니다!)
	// =========================================================
	fmt.Println("\n=== range 순회 ===")
	for fruit, p := range prices {
		fmt.Printf("%s: %d원\n", fruit, p)
	}

	// 키만 필요할 때
	fmt.Println("\n키만:")
	for fruit := range prices {
		fmt.Println("-", fruit)
	}

	// =========================================================
	// 맵으로 단어 빈도 세기 (실용적인 예제)
	// =========================================================
	fmt.Println("\n=== 단어 빈도 세기 ===")
	words := []string{"go", "is", "fun", "go", "is", "go"}
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++ // 없는 키는 제로값(0)에서 시작
	}
	fmt.Println("단어 빈도:", freq)
}
