package main

import (
	"fmt"
	"math"
)

// 인터페이스 정의: 메서드 시그니처만 선언
// 이 인터페이스를 "구현"하려면 두 메서드를 모두 가져야 합니다.
// implements 키워드 없이 자동으로 구현됩니다!
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle 구조체
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle 구조체
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Triangle 구조체
type Triangle struct {
	A, B, C float64 // 세 변의 길이
}

func (t Triangle) Area() float64 {
	// 헤론의 공식
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// Shape 인터페이스를 매개변수로 받는 함수 (다형성!)
// Circle이든 Rectangle이든 Shape를 구현하면 모두 받을 수 있습니다.
func printShapeInfo(s Shape) {
	fmt.Printf("  넓이: %.2f, 둘레: %.2f\n", s.Area(), s.Perimeter())
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

func main() {
	fmt.Println("=== 인터페이스 다형성 ===")

	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 3}
	t := Triangle{A: 3, B: 4, C: 5}

	fmt.Println("Circle:")
	printShapeInfo(c)

	fmt.Println("Rectangle:")
	printShapeInfo(r)

	fmt.Println("Triangle:")
	printShapeInfo(t)

	// Shape 인터페이스 슬라이스
	fmt.Println("\n=== Shape 슬라이스 ===")
	shapes := []Shape{c, r, t}
	fmt.Printf("전체 넓이 합계: %.2f\n", totalArea(shapes))

	// 타입 확인 (type assertion)
	fmt.Println("\n=== 타입 확인 (type assertion) ===")
	for _, s := range shapes {
		switch v := s.(type) {
		case Circle:
			fmt.Printf("원 (반지름=%.0f)\n", v.Radius)
		case Rectangle:
			fmt.Printf("직사각형 (%.0f×%.0f)\n", v.Width, v.Height)
		case Triangle:
			fmt.Printf("삼각형 (변: %.0f, %.0f, %.0f)\n", v.A, v.B, v.C)
		}
	}

	// 빈 인터페이스 (any): 모든 타입을 담을 수 있습니다
	fmt.Println("\n=== any (빈 인터페이스) ===")
	var anything interface{} // = any
	anything = 42
	fmt.Println("int:", anything)
	anything = "hello"
	fmt.Println("string:", anything)
	anything = true
	fmt.Println("bool:", anything)
}
