package main

import "fmt"

// 포인터를 사용하지 않는 함수: 값의 복사본이 전달됨
func doubleValue(n int) {
	n = n * 2 // 복사본만 수정됨, 원본은 변경 안됨
}

// 포인터를 사용하는 함수: 원본 변수를 직접 수정
func doublePointer(n *int) {
	*n = *n * 2 // *n: 포인터가 가리키는 실제 값
}

type Person struct {
	Name string
	Age  int
}

// 구조체를 포인터로 받아서 수정
func birthday(p *Person) {
	p.Age++ // (*p).Age++ 와 동일 (Go는 자동으로 역참조)
}

func main() {
	// =========================================================
	// 포인터 기초
	// & : 변수의 메모리 주소를 얻습니다
	// * : 포인터가 가리키는 값을 얻습니다 (역참조)
	// =========================================================
	fmt.Println("=== 포인터 기초 ===")
	x := 42
	p := &x // p는 x의 주소를 담은 *int 타입 포인터

	fmt.Printf("x의 값: %d\n", x)
	fmt.Printf("x의 주소: %v\n", &x)
	fmt.Printf("p (포인터 값): %v\n", p)    // 주소값
	fmt.Printf("*p (역참조): %d\n", *p) // 주소에 있는 값 = x의 값

	// 포인터로 원본 값 변경
	*p = 100
	fmt.Printf("*p = 100 후, x = %d\n", x) // x도 100이 됨

	// =========================================================
	// 값 전달 vs 포인터 전달
	// =========================================================
	fmt.Println("\n=== 값 전달 vs 포인터 전달 ===")
	num := 10
	doubleValue(num)
	fmt.Printf("doubleValue 후: %d (변화 없음)\n", num)

	doublePointer(&num) // &num으로 주소 전달
	fmt.Printf("doublePointer 후: %d (변경됨!)\n", num)

	// =========================================================
	// 구조체와 포인터
	// =========================================================
	fmt.Println("\n=== 구조체 포인터 ===")
	person := Person{Name: "홍길동", Age: 25}
	fmt.Printf("생일 전: %s %d살\n", person.Name, person.Age)

	birthday(&person)
	fmt.Printf("생일 후: %s %d살\n", person.Name, person.Age)

	// new(): 포인터를 바로 생성
	pp := new(Person)
	pp.Name = "김철수"
	pp.Age = 30
	fmt.Printf("new로 생성: %+v\n", *pp) // %+v: 필드 이름 포함 출력

	// =========================================================
	// nil 포인터 주의!
	// =========================================================
	fmt.Println("\n=== nil 포인터 주의 ===")
	var nilPtr *int
	fmt.Printf("nil 포인터: %v\n", nilPtr)

	// 아래 주석을 풀면 런타임 패닉 발생!
	// fmt.Println(*nilPtr) // panic: runtime error: invalid memory address

	// 항상 nil 체크 후 사용
	if nilPtr != nil {
		fmt.Println(*nilPtr)
	} else {
		fmt.Println("포인터가 nil입니다. 역참조 불가.")
	}
}
