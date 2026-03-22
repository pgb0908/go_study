package main

import "fmt"

// fmt.Stringer 인터페이스:
// String() string 메서드를 구현하면 fmt.Println이 자동으로 이 메서드를 사용합니다.
//
// type Stringer interface {
//     String() string
// }

type Color int

const (
	Red Color = iota
	Green
	Blue
)

// Color 타입에 String() 구현
func (c Color) String() string {
	switch c {
	case Red:
		return "빨강"
	case Green:
		return "초록"
	case Blue:
		return "파랑"
	default:
		return "알 수 없는 색"
	}
}

// =========================================================
// 구조체에 String() 구현
// =========================================================
type Point struct {
	X, Y float64
}

// String()이 없으면: {3.5 -2.1}
// String()이 있으면: (3.50, -2.10) 형식으로 출력
func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

type Student struct {
	Name  string
	Grade int
	Score float64
}

func (s Student) String() string {
	var grade string
	switch {
	case s.Score >= 90:
		grade = "A"
	case s.Score >= 80:
		grade = "B"
	case s.Score >= 70:
		grade = "C"
	default:
		grade = "D"
	}
	return fmt.Sprintf("[%d학년] %s - %.1f점 (%s등급)", s.Grade, s.Name, s.Score, grade)
}

func main() {
	// String()이 없으면 숫자로만 출력됨
	// String()을 구현하면 의미있는 텍스트로 출력
	fmt.Println("=== Color Stringer ===")
	c := Red
	fmt.Println(c)              // "빨강" (String() 자동 호출)
	fmt.Printf("색상: %v\n", c) // "색상: 빨강"
	fmt.Printf("색상: %s\n", c) // "색상: 빨강"

	colors := []Color{Red, Green, Blue}
	fmt.Println("모든 색상:", colors) // [빨강 초록 파랑]

	fmt.Println("\n=== Point Stringer ===")
	p := Point{X: 3.5, Y: -2.1}
	fmt.Println("점:", p)              // "(3.50, -2.10)"
	fmt.Printf("위치: %v\n", p)       // "(3.50, -2.10)"

	fmt.Println("\n=== Student Stringer ===")
	students := []Student{
		{Name: "홍길동", Grade: 3, Score: 92.5},
		{Name: "김철수", Grade: 2, Score: 78.0},
		{Name: "이영희", Grade: 1, Score: 85.5},
	}
	for _, s := range students {
		fmt.Println(s)
	}
}
