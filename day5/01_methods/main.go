package main

import (
	"fmt"
	"math"
)

// 구조체 정의
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// =========================================================
// 값 리시버 (Value Receiver): (c Circle)
// 구조체의 복사본으로 동작합니다.
// 메서드 안에서 수정해도 원본에 영향 없음.
// 읽기 전용 작업에 사용합니다.
// =========================================================

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// =========================================================
// 포인터 리시버 (Pointer Receiver): (c *Circle)
// 원본 구조체를 직접 수정합니다.
// 구조체 필드를 변경하는 메서드에 사용합니다.
// =========================================================

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor // 원본의 Radius가 변경됨
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(반지름=%.2f)", c.Radius)
}

func main() {
	c := Circle{Radius: 5}
	fmt.Println("=== Circle 메서드 ===")
	fmt.Printf("원: %s\n", c.String())
	fmt.Printf("넓이: %.2f\n", c.Area())
	fmt.Printf("둘레: %.2f\n", c.Perimeter())

	// 포인터 리시버 메서드 호출
	// Go는 자동으로 &c를 전달합니다 (c.Scale은 (&c).Scale과 동일)
	c.Scale(2)
	fmt.Printf("2배 확대 후: %s\n", c.String())

	fmt.Println("\n=== Rectangle 메서드 ===")
	r := Rectangle{Width: 4, Height: 3}
	fmt.Printf("직사각형: 너비=%.0f, 높이=%.0f\n", r.Width, r.Height)
	fmt.Printf("넓이: %.2f\n", r.Area())
	fmt.Printf("둘레: %.2f\n", r.Perimeter())

	// =========================================================
	// 값 리시버 vs 포인터 리시버 차이 확인
	// =========================================================
	fmt.Println("\n=== 값 리시버 vs 포인터 리시버 ===")
	c1 := Circle{Radius: 10}
	c2 := c1 // 복사

	c1.Scale(2) // 포인터 리시버: c1만 변경됨
	fmt.Printf("c1: %s\n", c1.String()) // 반지름 20
	fmt.Printf("c2: %s (복사본, 영향 없음)\n", c2.String()) // 반지름 10
}
