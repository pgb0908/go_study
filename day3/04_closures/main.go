package main

import "fmt"

// 클로저: 함수가 자신을 둘러싼 환경의 변수를 "포획(capture)"합니다.
// makeCounter는 새로운 카운터 함수를 반환합니다.
func makeCounter() func() int {
	count := 0 // 이 변수는 반환된 함수가 살아있는 한 유지됩니다
	return func() int {
		count++ // 외부 변수 count를 포획해서 계속 참조
		return count
	}
}

// 클로저로 "곱하기 N" 함수 생성기
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor // factor를 포획
	}
}

func main() {
	// =========================================================
	// 클로저: 독립적인 상태 유지
	// =========================================================
	fmt.Println("=== 클로저 - 카운터 ===")
	counter1 := makeCounter()
	counter2 := makeCounter() // 완전히 독립적인 카운터

	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2
	fmt.Println(counter1()) // 3
	fmt.Println(counter2()) // 1 (counter1과 별개!)
	fmt.Println(counter2()) // 2

	// =========================================================
	// 클로저로 함수 생성기
	// =========================================================
	fmt.Println("\n=== 클로저 - 함수 생성기 ===")
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Printf("double(5) = %d\n", double(5))
	fmt.Printf("triple(5) = %d\n", triple(5))

	// =========================================================
	// defer: 함수가 끝날 때 실행되는 코드
	// 주로 자원 정리(파일 닫기, DB 연결 해제 등)에 사용
	// =========================================================
	fmt.Println("\n=== defer ===")
	deferExample()

	// =========================================================
	// defer는 LIFO 순서 (마지막에 선언된 것이 먼저 실행)
	// =========================================================
	fmt.Println("\n=== defer 실행 순서 (LIFO) ===")
	deferOrder()
}

func deferExample() {
	fmt.Println("함수 시작")
	defer fmt.Println("defer: 함수 끝날 때 실행") // 나중에 실행
	fmt.Println("함수 중간")
	// deferExample이 끝나면 defer가 실행됨
}

func deferOrder() {
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("defer %d 실행\n", i)
	}
	fmt.Println("반복문 완료")
	// 출력 순서: 반복문 완료 → defer 3 → defer 2 → defer 1 (LIFO)
}
