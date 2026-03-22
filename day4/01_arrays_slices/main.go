package main

import "fmt"

func main() {
	// =========================================================
	// 배열 (Array): 크기가 고정됩니다
	// =========================================================
	fmt.Println("=== 배열 ===")
	var arr [3]int // [3]int = 크기 3인 int 배열, 제로값으로 초기화
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("배열:", arr)
	fmt.Printf("길이: %d\n", len(arr))

	// 리터럴로 초기화
	fruits := [3]string{"사과", "바나나", "딸기"}
	fmt.Println("과일:", fruits)

	// =========================================================
	// 슬라이스 (Slice): 크기가 동적으로 변합니다
	// 실제로는 배열보다 슬라이스를 훨씬 많이 씁니다!
	// =========================================================
	fmt.Println("\n=== 슬라이스 ===")
	s := []int{1, 2, 3} // []int (대괄호 안에 숫자 없음 = 슬라이스)
	fmt.Printf("슬라이스: %v, 길이(len): %d, 용량(cap): %d\n", s, len(s), cap(s))

	// append: 요소를 추가합니다 (필요 시 내부 배열이 자동으로 늘어남)
	s = append(s, 4)
	s = append(s, 5, 6, 7) // 여러 개 한번에 추가
	fmt.Printf("append 후: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// make: 길이와 용량을 지정해서 슬라이스 생성
	s2 := make([]int, 3, 5) // 길이 3, 용량 5 (미리 공간 확보)
	fmt.Printf("make: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	// =========================================================
	// 슬라이싱: 슬라이스의 일부를 잘라냅니다
	// s[시작:끝] - 끝 인덱스는 포함되지 않습니다
	// =========================================================
	fmt.Println("\n=== 슬라이싱 ===")
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("원본:", nums)
	fmt.Println("nums[2:5]:", nums[2:5]) // 인덱스 2,3,4
	fmt.Println("nums[:3]:", nums[:3])   // 처음부터 인덱스 0,1,2
	fmt.Println("nums[7:]:", nums[7:])   // 인덱스 7부터 끝까지

	// =========================================================
	// copy: 슬라이스 완전 복사 (슬라이스 직접 대입은 같은 배열을 공유!)
	// =========================================================
	fmt.Println("\n=== copy ===")
	original := []int{1, 2, 3}
	copied := make([]int, len(original))
	copy(copied, original)

	copied[0] = 999
	fmt.Println("original:", original) // [1 2 3] - 영향 없음
	fmt.Println("copied:", copied)     // [999 2 3]

	// 주의: 대입으로 복사하면 같은 배열을 참조!
	shared := original
	shared[0] = 999
	fmt.Println("\n주의! 대입 후 original:", original) // [999 2 3] - 변경됨!
}
