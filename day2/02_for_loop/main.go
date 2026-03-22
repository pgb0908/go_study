package main

import "fmt"

func main() {
	// =========================================================
	// for 방법 1: C 스타일 (초기화; 조건; 후처리)
	// =========================================================
	fmt.Println("=== C 스타일 for ===")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// =========================================================
	// for 방법 2: 조건만 있는 for (다른 언어의 while과 동일)
	// =========================================================
	fmt.Println("\n=== while 스타일 for ===")
	n := 1
	for n < 100 {
		n *= 2 // 2배씩 증가
	}
	fmt.Printf("100보다 큰 최소 2의 거듭제곱: %d\n", n)

	// =========================================================
	// for 방법 3: 무한 루프 (break로 탈출)
	// =========================================================
	fmt.Println("\n=== 무한 루프 + break ===")
	count := 0
	for {
		count++
		if count >= 3 {
			break // 루프 탈출
		}
		fmt.Println("반복 중...", count)
	}
	fmt.Println("루프 종료")

	// continue: 이번 반복을 건너뛰고 다음으로
	fmt.Println("\n=== continue (홀수만 출력) ===")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // 짝수는 건너뜀
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// =========================================================
	// for 방법 4: range (슬라이스, 문자열, 맵 순회)
	// =========================================================
	fmt.Println("\n=== range로 슬라이스 순회 ===")
	fruits := []string{"사과", "바나나", "딸기"}
	for i, fruit := range fruits {
		fmt.Printf("[%d] %s\n", i, fruit)
	}

	// 인덱스가 필요 없을 때: _ (blank identifier)로 버림
	fmt.Println("\n=== 값만 사용 ===")
	sum := 0
	numbers := []int{10, 20, 30, 40, 50}
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("합계: %d\n", sum)

	// 문자열도 range로 순회 가능 (문자 단위)
	fmt.Println("\n=== 문자열 range ===")
	for i, ch := range "Hello" {
		fmt.Printf("[%d] %c\n", i, ch)
	}
}
