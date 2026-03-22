package main

import "fmt"

func main() {
	// =========================================================
	// 변수 선언 방법 1: var + 타입 명시
	// =========================================================
	var name string = "홍길동"
	var age int = 25
	var height float64 = 175.5
	var isStudent bool = true

	fmt.Println("=== 방법 1: var + 타입 명시 ===")
	fmt.Printf("이름: %s\n", name)
	fmt.Printf("나이: %d\n", age)
	fmt.Printf("키: %.1f cm\n", height) // .1f = 소수점 1자리
	fmt.Printf("학생 여부: %v\n", isStudent)

	// =========================================================
	// 변수 선언 방법 2: var + 타입 추론 (컴파일러가 타입을 결정)
	// =========================================================
	var city = "서울"       // string으로 추론
	var temperature = 36.5 // float64로 추론

	fmt.Println("\n=== 방법 2: var + 타입 추론 ===")
	fmt.Printf("도시: %s\n", city)
	fmt.Printf("체온: %.1f°C\n", temperature)

	// =========================================================
	// 변수 선언 방법 3: 짧은 선언 := (함수 안에서만 사용 가능!)
	// =========================================================
	job := "개발자"  // string
	salary := 5000 // int

	fmt.Println("\n=== 방법 3: 짧은 선언 := ===")
	fmt.Printf("직업: %s\n", job)
	fmt.Printf("월급: %d만원\n", salary)

	// =========================================================
	// Go의 제로값 (Zero Value)
	// 값을 지정하지 않으면 타입에 따라 기본값이 자동 설정됩니다
	// =========================================================
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Println("\n=== 제로값 (기본값) ===")
	fmt.Printf("int 제로값: %d\n", zeroInt)
	fmt.Printf("float64 제로값: %f\n", zeroFloat)
	fmt.Printf("string 제로값: '%s'\n", zeroString)
	fmt.Printf("bool 제로값: %v\n", zeroBool)

	// =========================================================
	// 여러 변수 동시 선언
	// =========================================================
	var (
		x = 10
		y = 20
		z = 30
	)
	fmt.Printf("\nx=%d, y=%d, z=%d, 합계=%d\n", x, y, z, x+y+z)
}
