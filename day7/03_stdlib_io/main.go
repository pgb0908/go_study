package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// =========================================================
	// os.Args: 커맨드라인 인자
	// 실행: go run ./day7/03_stdlib_io/ 홍길동 25
	// =========================================================
	fmt.Println("=== os.Args (커맨드라인 인자) ===")
	fmt.Printf("인자 개수: %d\n", len(os.Args))
	for i, arg := range os.Args {
		fmt.Printf("Args[%d] = %s\n", i, arg)
	}
	// Args[0]은 프로그램 경로, 실제 인자는 Args[1]부터

	// =========================================================
	// 파일 쓰기
	// =========================================================
	fmt.Println("\n=== 파일 쓰기 ===")
	content := "첫 번째 줄\n두 번째 줄\n세 번째 줄\n"

	// os.WriteFile: 간단한 파일 쓰기 (Go 1.16+)
	err := os.WriteFile("/tmp/test_go_study.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("쓰기 에러:", err)
		return
	}
	fmt.Println("파일 쓰기 완료: /tmp/test_go_study.txt")

	// =========================================================
	// 파일 읽기 - 방법 1: 전체 한번에 읽기
	// =========================================================
	fmt.Println("\n=== 파일 읽기 (전체) ===")
	data, err := os.ReadFile("/tmp/test_go_study.txt")
	if err != nil {
		fmt.Println("읽기 에러:", err)
		return
	}
	fmt.Print(string(data))

	// =========================================================
	// 파일 읽기 - 방법 2: 줄 단위로 읽기 (bufio.Scanner)
	// 큰 파일에 적합합니다
	// =========================================================
	fmt.Println("=== 파일 읽기 (줄 단위) ===")
	f, err := os.Open("/tmp/test_go_study.txt")
	if err != nil {
		fmt.Println("열기 에러:", err)
		return
	}
	defer f.Close() // 함수 종료 시 파일 닫기 (defer!)

	scanner := bufio.NewScanner(f)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("줄 %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	// =========================================================
	// 표준 입력에서 줄 단위로 읽기
	// =========================================================
	fmt.Println("\n=== 표준 입력 읽기 ===")
	fmt.Println("문장을 입력하세요 (단어 수를 세어 드립니다):")
	fmt.Println("(빈 줄로 종료)")

	inputScanner := bufio.NewScanner(os.Stdin)
	totalWords := 0
	for inputScanner.Scan() {
		line := inputScanner.Text()
		if line == "" {
			break
		}
		words := strings.Fields(line)
		totalWords += len(words)
		fmt.Printf("  → %d개 단어\n", len(words))
	}
	fmt.Printf("총 단어 수: %d\n", totalWords)

	// 파일 삭제 (정리)
	os.Remove("/tmp/test_go_study.txt")
}
