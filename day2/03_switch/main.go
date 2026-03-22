package main

import "fmt"

func main() {
	// =========================================================
	// 기본 switch
	// Go의 switch는 각 case 끝에 자동으로 break됩니다!
	// break를 직접 쓸 필요가 없습니다.
	// =========================================================
	fmt.Println("=== 기본 switch ===")
	day := "월요일"

	switch day {
	case "토요일", "일요일": // 여러 값을 하나의 case에 묶을 수 있습니다
		fmt.Println("주말입니다!")
	case "월요일":
		fmt.Println("한 주의 시작!")
	case "금요일":
		fmt.Println("곧 주말!")
	default:
		fmt.Println("평일입니다")
	}

	// =========================================================
	// 표현식 없는 switch (if-else chain 대체)
	// =========================================================
	fmt.Println("\n=== 표현식 없는 switch ===")
	score := 75

	switch {
	case score >= 90:
		fmt.Println("A등급: 우수")
	case score >= 80:
		fmt.Println("B등급: 양호")
	case score >= 70:
		fmt.Println("C등급: 보통")
	default:
		fmt.Println("D등급: 미흡")
	}

	// =========================================================
	// switch 초기화 구문 (if처럼 사용 가능)
	// =========================================================
	fmt.Println("\n=== switch 초기화 구문 ===")
	switch remainder := score % 10; {
	case remainder >= 7:
		fmt.Printf("소수점 올림 (나머지: %d)\n", remainder)
	default:
		fmt.Printf("소수점 내림 (나머지: %d)\n", remainder)
	}

	// =========================================================
	// fallthrough: 다음 case도 실행하고 싶을 때
	// (명시적으로 선언해야만 다음 case로 넘어감)
	// =========================================================
	fmt.Println("\n=== fallthrough ===")
	num := 2
	switch num {
	case 1:
		fmt.Println("1입니다")
		fallthrough
	case 2:
		fmt.Println("1 또는 2입니다") // num=2이지만 이 줄이 실행됨
		fallthrough
	case 3:
		fmt.Println("1, 2 또는 3입니다") // fallthrough로 여기도 실행됨
	case 4:
		fmt.Println("4입니다") // 여기는 실행 안됨
	}
}
