package homework

import (
	"math"
	"testing"
)

// Q1. 테이블 주도 테스트 - 아래 케이스를 채워보세요!
func TestFibonacci(t *testing.T) {
	cases := []struct {
		n, want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{5, 5},
		{10, 55},
		// TODO: 케이스를 더 추가해보세요
	}
	for _, c := range cases {
		got := Fibonacci(c.n)
		if got != c.want {
			t.Errorf("Q1) Fibonacci(%d): got %d, want %d", c.n, got, c.want)
		}
	}
}

func TestProcessString(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"  Hello World  ", "hello_world"},
		{"Go Language", "go_language"},
		{"  already  ", "already"},
		{"NoSpaces", "nospaces"},
	}
	for _, c := range cases {
		got := ProcessString(c.input)
		if got != c.want {
			t.Errorf("Q2) ProcessString(%q): got %q, want %q", c.input, got, c.want)
		}
	}
}

func TestGradeReport(t *testing.T) {
	// 일반 케이스
	scores := map[string]int{"Alice": 90, "Bob": 75, "Carol": 88}
	name, avg := GradeReport(scores)
	if name != "Alice" {
		t.Errorf("Q3) 최고점 학생: got %q, want %q", name, "Alice")
	}
	wantAvg := float64(90+75+88) / 3.0
	if math.Abs(avg-wantAvg) > 0.01 {
		t.Errorf("Q3) 평균: got %v, want %v", avg, wantAvg)
	}

	// 빈 맵
	name2, avg2 := GradeReport(map[string]int{})
	if name2 != "" || avg2 != 0.0 {
		t.Errorf("Q3) 빈 맵: got (%q,%v), want (\"\",0.0)", name2, avg2)
	}

	// 단일 학생
	name3, avg3 := GradeReport(map[string]int{"Solo": 100})
	if name3 != "Solo" || avg3 != 100 {
		t.Errorf("Q3) 단일 학생: got (%q,%v), want (\"Solo\",100)", name3, avg3)
	}
}
