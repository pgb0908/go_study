package main

import (
	"errors"
	"fmt"
)

// =========================================================
// 방법 1: errors.New - 간단한 에러 메시지
// =========================================================
var ErrDivByZero = errors.New("0으로 나눌 수 없습니다")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// =========================================================
// 방법 2: fmt.Errorf - 값을 포함한 에러 메시지
// %w: 에러를 "래핑(wrapping)"합니다 (errors.Is로 비교 가능)
// =========================================================
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("음수(%v)의 제곱근은 계산할 수 없습니다", n)
	}
	// 뉴턴법으로 제곱근 계산
	x := n
	for i := 0; i < 100; i++ {
		x = (x + n/x) / 2
	}
	return x, nil
}

// =========================================================
// 방법 3: 커스텀 에러 타입
// error 인터페이스: Error() string 메서드만 구현하면 됩니다
// =========================================================
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("유효성 검사 실패 - %s(%v): %s", e.Field, e.Value, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "나이는 0 이상이어야 합니다",
		}
	}
	if age > 150 {
		return &ValidationError{
			Field:   "age",
			Value:   age,
			Message: "나이는 150 이하이어야 합니다",
		}
	}
	return nil
}

func main() {
	// =========================================================
	// 기본 에러 처리
	// =========================================================
	fmt.Println("=== 기본 에러 처리 ===")
	if result, err := divide(10, 3); err != nil {
		fmt.Println("에러:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	if _, err := divide(10, 0); err != nil {
		fmt.Println("에러:", err)
	}

	// =========================================================
	// errors.Is: 특정 에러인지 확인
	// =========================================================
	fmt.Println("\n=== errors.Is ===")
	_, err := divide(5, 0)
	if errors.Is(err, ErrDivByZero) {
		fmt.Println("0으로 나누기 에러 감지!")
	}

	// =========================================================
	// 커스텀 에러 타입 + errors.As
	// =========================================================
	fmt.Println("\n=== 커스텀 에러 ===")
	ages := []int{25, -5, 200}
	for _, age := range ages {
		if err := validateAge(age); err != nil {
			// errors.As: 특정 타입의 에러인지 확인하고 변환
			var ve *ValidationError
			if errors.As(err, &ve) {
				fmt.Printf("검증 에러 - 필드: %s, 값: %v\n", ve.Field, ve.Value)
			}
		} else {
			fmt.Printf("나이 %d: 유효합니다\n", age)
		}
	}

	// =========================================================
	// 에러 무시 (비권장! 하지만 알아야 함)
	// =========================================================
	fmt.Println("\n=== 에러 처리 패턴 ===")
	result, _ := sqrt(16) // 에러를 _ 로 무시 (확실히 에러 없을 때만 사용)
	fmt.Printf("√16 = %.4f\n", result)

	if _, err := sqrt(-4); err != nil {
		fmt.Println("에러:", err)
	}
}
