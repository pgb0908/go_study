package main

import (
	"testing"
)

// =========================================================
// 테이블 주도 테스트 (Table-Driven Tests)
// Go에서 가장 권장하는 테스트 패턴입니다.
// =========================================================

func TestAdd(t *testing.T) {
	// 테스트 케이스 테이블 정의
	tests := []struct {
		name     string // 테스트 케이스 이름
		a, b     int
		expected int
	}{
		{"양수 덧셈", 2, 3, 5},
		{"음수 포함", -1, 1, 0},
		{"둘 다 음수", -3, -4, -7},
		{"영 포함", 0, 5, 5},
		{"큰 수", 100, 200, 300},
	}

	for _, tt := range tests {
		// t.Run: 서브테스트 실행 (케이스별로 독립 실행)
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				// t.Errorf: 실패 기록 후 계속 진행
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"기본 빼기", 10, 3, 7},
		{"음수 결과", 3, 10, -7},
		{"같은 수", 5, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	// 성공 케이스
	t.Run("정상 나눗셈", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("예상치 못한 에러: %v", err) // t.Fatal: 즉시 중단
		}
		if got != 5.0 {
			t.Errorf("Divide(10, 2) = %f; want 5.0", got)
		}
	})

	// 에러 케이스
	t.Run("0으로 나누기", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			// 에러가 발생해야 하는데 발생하지 않으면 실패
			t.Error("0으로 나누기에서 에러가 반환되어야 합니다")
		}
	})
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{17, true},
		{20, false},
		{97, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := IsPrime(tt.n)
			if got != tt.expected {
				t.Errorf("IsPrime(%d) = %v; want %v", tt.n, got, tt.expected)
			}
		})
	}
}
