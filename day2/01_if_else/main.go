package main

import "fmt"

func main() {
	// =========================================================
	// 기본 if-else
	// =========================================================
	score := 85

	if score >= 90 {
		fmt.Println("A등급")
	} else if score >= 80 {
		fmt.Println("B등급")
	} else if score >= 70 {
		fmt.Println("C등급")
	} else {
		fmt.Println("D등급")
	}

	// =========================================================
	// Go만의 특징: if 초기화 구문
	// 조건식 앞에 짧은 문장을 실행할 수 있습니다.
	// 변수의 범위(scope)가 if 블록 안으로 제한되어 깔끔합니다.
	// =========================================================
	fmt.Println("\n=== if 초기화 구문 ===")

	// remainder는 if-else 블록 밖에서는 사용할 수 없습니다
	if remainder := score % 10; remainder >= 5 {
		fmt.Printf("점수 %d: 올림 처리 (나머지 %d)\n", score, remainder)
	} else {
		fmt.Printf("점수 %d: 내림 처리 (나머지 %d)\n", score, remainder)
	}

	// 이 패턴은 에러 처리에서 매우 자주 씁니다 (Day3에서 다시 등장)
	// if err := someFunc(); err != nil { ... }

	// =========================================================
	// 논리 연산자
	// =========================================================
	fmt.Println("\n=== 논리 연산자 ===")
	age := 20
	hasID := true

	if age >= 18 && hasID {
		fmt.Println("입장 가능합니다")
	}

	isWeekend := false
	isHoliday := true

	if isWeekend || isHoliday {
		fmt.Println("오늘은 쉬는 날입니다")
	}

	isRaining := false
	if !isRaining {
		fmt.Println("우산이 필요 없습니다")
	}
}
