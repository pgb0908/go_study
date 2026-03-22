package main

import "fmt"

// 구조체 (Struct): 여러 필드를 하나로 묶는 타입
// Go에는 클래스(class)가 없습니다. 구조체 + 메서드로 대체합니다.
type Person struct {
	Name string
	Age  int
	City string
}

// 구조체도 다른 구조체를 포함할 수 있습니다 (임베딩 맛보기)
type Address struct {
	Street string
	City   string
}

type Employee struct {
	Person  // 임베딩: Person의 필드를 직접 접근 가능
	Address // 임베딩
	Company string
	Salary  int
}

func main() {
	// =========================================================
	// 구조체 생성과 초기화
	// =========================================================
	fmt.Println("=== 구조체 생성 ===")

	// 방법 1: 필드 이름으로 초기화 (권장)
	p1 := Person{
		Name: "홍길동",
		Age:  25,
		City: "서울",
	}

	// 방법 2: 순서대로 초기화 (필드 순서를 정확히 알아야 함 - 비권장)
	p2 := Person{"김철수", 30, "부산"}

	// 방법 3: 제로값으로 시작 후 필드 설정
	var p3 Person
	p3.Name = "이영희"
	p3.Age = 22

	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
	fmt.Println("p3:", p3)

	// =========================================================
	// 필드 접근
	// =========================================================
	fmt.Println("\n=== 필드 접근 ===")
	fmt.Printf("이름: %s, 나이: %d살, 도시: %s\n", p1.Name, p1.Age, p1.City)
	p1.Age = 26 // 필드 수정
	fmt.Printf("생일 후 나이: %d살\n", p1.Age)

	// =========================================================
	// 익명 구조체: 타입 이름 없이 바로 정의
	// 일회성으로 사용할 때 편리합니다
	// =========================================================
	fmt.Println("\n=== 익명 구조체 ===")
	config := struct {
		Host string
		Port int
	}{
		Host: "localhost",
		Port: 8080,
	}
	fmt.Printf("서버: %s:%d\n", config.Host, config.Port)

	// =========================================================
	// 구조체 슬라이스
	// =========================================================
	fmt.Println("\n=== 구조체 슬라이스 ===")
	people := []Person{
		{Name: "철수", Age: 20, City: "서울"},
		{Name: "영희", Age: 25, City: "부산"},
		{Name: "민수", Age: 18, City: "대구"},
	}
	for _, p := range people {
		fmt.Printf("%-5s %d살 (%s)\n", p.Name, p.Age, p.City)
	}

	// =========================================================
	// 구조체 임베딩
	// =========================================================
	fmt.Println("\n=== 구조체 임베딩 ===")
	emp := Employee{
		Person:  Person{Name: "박지성", Age: 35, City: "서울"},
		Address: Address{Street: "테헤란로 123", City: "서울"},
		Company: "Go Corp",
		Salary:  5000,
	}
	// 임베딩된 필드는 직접 접근 가능
	fmt.Printf("이름: %s (임베딩 직접 접근)\n", emp.Name) // emp.Person.Name과 동일
	fmt.Printf("회사: %s, 월급: %d만원\n", emp.Company, emp.Salary)
}
