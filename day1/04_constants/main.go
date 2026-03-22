package main

import "fmt"

// 상수(const): 한번 선언하면 값을 바꿀 수 없습니다.
// 변수와 달리 패키지 레벨에서도 선언 가능합니다.
const Pi = 3.14159
const AppName = "Go학습앱"

// iota: 상수 그룹에서 0부터 자동으로 증가하는 특별한 값
// 열거형(enum) 패턴을 만들 때 사용합니다.
const (
	Sunday    = iota // 0
	Monday           // 1 (자동으로 이전 값 + 1)
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// iota에 수식을 적용할 수도 있습니다
const (
	_  = iota // 0은 버림 (blank identifier _)
	KB = 1 << (10 * iota) // 1 << 10 = 1024
	MB                    // 1 << 20 = 1,048,576
	GB                    // 1 << 30 = 1,073,741,824
)

func main() {
	fmt.Println("=== 기본 상수 ===")
	fmt.Printf("원주율: %f\n", Pi)
	fmt.Printf("앱 이름: %s\n", AppName)

	fmt.Println("\n=== 요일 상수 (iota) ===")
	fmt.Printf("일요일=%d, 월요일=%d, 금요일=%d\n", Sunday, Monday, Friday)

	today := Wednesday
	if today == Wednesday {
		fmt.Println("오늘은 수요일입니다!")
	}

	fmt.Println("\n=== 파일 크기 상수 ===")
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)

	// 상수 vs 변수: 상수는 변경 불가
	// Pi = 3.14  // 이 줄의 주석을 풀면 컴파일 에러 발생!

	fmt.Println("\n=== 상수로 계산하기 ===")
	radius := 5.0
	area := Pi * radius * radius
	fmt.Printf("반지름 %.0f인 원의 넓이: %.2f\n", radius, area)
}
