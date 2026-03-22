package main

import "fmt"

func main() {
	// fmt.Scan 계열 함수로 사용자 입력을 받습니다.
	// 중요: 변수 앞에 & (주소 연산자)를 붙여야 합니다!
	// &name 은 "name 변수가 저장된 메모리 주소"를 의미합니다.

	var name string
	fmt.Print("이름을 입력하세요: ")
	fmt.Scan(&name) // 공백/줄바꿈까지 읽음

	var age int
	fmt.Print("나이를 입력하세요: ")
	fmt.Scan(&age)

	// Println vs Printf 비교
	// Println: 자동 줄바꿈, 형식 없이 출력
	// Printf: 형식 문자열 사용, \n 직접 입력 필요
	fmt.Printf("\n안녕하세요, %d살 %s님!\n", age, name)

	// Sprintf: 출력하지 않고 문자열로 반환
	greeting := fmt.Sprintf("반갑습니다, %s님!", name)
	fmt.Println(greeting)
}
