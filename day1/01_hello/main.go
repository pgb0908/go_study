// Go의 모든 파일은 반드시 패키지 선언으로 시작합니다.
// "main" 패키지는 특별합니다 - 실행 가능한 프로그램의 진입점입니다.
package main

// import: 다른 패키지의 코드를 가져옵니다.
// "fmt"는 Go 표준 라이브러리의 형식 입출력 패키지입니다.
import "fmt"

// func main(): 프로그램이 시작될 때 가장 먼저 실행되는 함수입니다.
// Java의 public static void main(String[] args)와 같은 역할이지만 훨씬 간결합니다.
func main() {
	// Println: Print + line의 줄임말. 출력 후 자동으로 줄바꿈합니다.
	fmt.Println("Hello, World!")

	// Printf: 형식 문자열로 출력합니다.
	// %s = 문자열, %d = 정수, %f = 실수, %v = 모든 타입 (범용)
	fmt.Printf("Hello, %s!\n", "Go")

	// \n은 줄바꿈 문자입니다. Printf는 자동 줄바꿈이 없어서 직접 넣어야 합니다.
	fmt.Printf("1 + 1 = %d\n", 1+1)
}
