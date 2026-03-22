package main

import "fmt"

// 기본 함수: func 이름(매개변수 타입) 반환타입
func add(a int, b int) int {
	return a + b
}

// 같은 타입의 매개변수는 타입을 한 번만 쓸 수 있습니다
func multiply(a, b int) int {
	return a * b
}

// 반환값이 없는 함수
func greet(name string) {
	fmt.Printf("안녕하세요, %s님!\n", name)
}

// 재귀 함수: 자기 자신을 호출하는 함수
// 팩토리얼: n! = n × (n-1) × ... × 1
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1) // 자기 자신 호출
}

func main() {
	fmt.Println("=== 기본 함수 ===")
	fmt.Printf("3 + 4 = %d\n", add(3, 4))
	fmt.Printf("3 × 4 = %d\n", multiply(3, 4))
	greet("홍길동")

	fmt.Println("\n=== 재귀 함수 (팩토리얼) ===")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}

	// =========================================================
	// 함수를 변수에 저장 (함수는 1급 시민입니다)
	// =========================================================
	fmt.Println("\n=== 함수를 변수에 저장 ===")

	// 함수 타입: func(매개변수타입) 반환타입
	var operation func(int, int) int

	operation = add
	fmt.Printf("더하기: %d\n", operation(10, 5))

	operation = multiply
	fmt.Printf("곱하기: %d\n", operation(10, 5))

	// 익명 함수 (이름 없이 바로 정의)
	square := func(n int) int {
		return n * n
	}
	fmt.Printf("5의 제곱: %d\n", square(5))

	// 즉시 실행 함수 (선언하자마자 바로 호출)
	result := func(x, y int) int {
		return x + y
	}(3, 7) // 바로 (3, 7)로 호출
	fmt.Printf("즉시 실행 결과: %d\n", result)
}
