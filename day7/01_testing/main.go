package main

import (
	"errors"
	"fmt"
)

// 테스트할 함수들
// main_test.go 파일에서 이 함수들을 테스트합니다.

// Add: 두 정수를 더합니다
func Add(a, b int) int {
	return a + b
}

// Subtract: 두 정수를 뺍니다
func Subtract(a, b int) int {
	return a - b
}

// Divide: 나눗셈 결과와 에러를 반환합니다
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0으로 나눌 수 없습니다")
	}
	return a / b, nil
}

// IsPrime: 소수 여부를 확인합니다
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("=== 함수 직접 실행 ===")
	fmt.Printf("Add(3, 4) = %d\n", Add(3, 4))
	fmt.Printf("Subtract(10, 3) = %d\n", Subtract(10, 3))

	result, err := Divide(10, 3)
	if err != nil {
		fmt.Println("에러:", err)
	} else {
		fmt.Printf("Divide(10, 3) = %.4f\n", result)
	}

	fmt.Println("\n=== 소수 목록 ===")
	for i := 2; i <= 20; i++ {
		if IsPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	fmt.Println("\n테스트 실행: go test ./day7/01_testing/")
	fmt.Println("상세 출력: go test -v ./day7/01_testing/")
}
