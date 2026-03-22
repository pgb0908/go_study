package main

import (
	"errors"
	"fmt"
)

// 다중 반환값: Go의 가장 강력한 특성 중 하나
// 에러 처리 패턴의 기초입니다.

// 두 값을 반환하는 함수
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max // 두 값을 동시에 반환
}

// 에러도 함께 반환하는 패턴 (Go에서 가장 흔한 패턴!)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// errors.New로 에러 값을 만듭니다
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil // 성공 시 error는 nil (없음)
}

// 명명된 반환값 (named return): 반환값에 이름을 붙일 수 있습니다
func rectangle(width, height float64) (area float64, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // 이름이 있으면 빈 return으로도 반환 가능 (naked return)
}

func main() {
	// 두 값 동시에 받기
	fmt.Println("=== 다중 반환값 ===")
	numbers := []int{5, 2, 8, 1, 9, 3}
	min, max := minMax(numbers)
	fmt.Printf("최솟값: %d, 최댓값: %d\n", min, max)

	// 하나만 필요할 때는 _ (blank identifier)로 버림
	_, onlyMax := minMax(numbers)
	fmt.Printf("최댓값만: %d\n", onlyMax)

	// =========================================================
	// 에러 처리 패턴 (Go에서 가장 중요한 패턴!)
	// =========================================================
	fmt.Println("\n=== 에러 처리 패턴 ===")

	// 성공 케이스
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("에러:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	// 실패 케이스
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("에러:", err) // 에러 출력
	} else {
		fmt.Printf("결과: %.4f\n", result)
	}

	// =========================================================
	// 명명된 반환값
	// =========================================================
	fmt.Println("\n=== 명명된 반환값 ===")
	a, p := rectangle(5, 3)
	fmt.Printf("너비=5, 높이=3: 넓이=%.1f, 둘레=%.1f\n", a, p)
}
