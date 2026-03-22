package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// =========================================================
	// strings 패키지: 문자열 처리
	// =========================================================
	fmt.Println("=== strings 패키지 ===")

	s := "Hello, Go World!"

	// 검색
	fmt.Println("Contains('Go'):", strings.Contains(s, "Go"))
	fmt.Println("HasPrefix('Hello'):", strings.HasPrefix(s, "Hello"))
	fmt.Println("HasSuffix('!'):", strings.HasSuffix(s, "!"))
	fmt.Println("Index('Go'):", strings.Index(s, "Go"))
	fmt.Println("Count('o'):", strings.Count(s, "o"))

	// 변환
	fmt.Println("\nToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("Replace:", strings.Replace(s, "Go", "Golang", 1))
	fmt.Println("ReplaceAll:", strings.ReplaceAll(s, "o", "0"))

	// 분리/결합
	fmt.Println("\n=== Split/Join ===")
	csv := "홍길동,25,서울,개발자"
	parts := strings.Split(csv, ",")
	fmt.Println("Split:", parts)
	fmt.Printf("parts[0]=%s, parts[1]=%s\n", parts[0], parts[1])

	joined := strings.Join(parts, " | ")
	fmt.Println("Join:", joined)

	// Fields: 공백(스페이스, 탭, 줄바꿈)으로 분리
	sentence := "  Go   is   awesome  "
	words := strings.Fields(sentence)
	fmt.Println("Fields:", words)

	// 공백 제거
	fmt.Println("\n=== Trim ===")
	padded := "   hello   "
	fmt.Printf("원본: '%s'\n", padded)
	fmt.Printf("TrimSpace: '%s'\n", strings.TrimSpace(padded))
	fmt.Printf("Trim('-', '---hello---'): '%s'\n", strings.Trim("---hello---", "-"))
	fmt.Printf("TrimLeft: '%s'\n", strings.TrimLeft(padded, " "))

	// strings.Builder: 문자열을 효율적으로 조립
	fmt.Println("\n=== strings.Builder ===")
	var sb strings.Builder
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(&sb, "항목%d", i)
		if i < 5 {
			sb.WriteString(", ")
		}
	}
	fmt.Println(sb.String())

	// =========================================================
	// strconv 패키지: 타입 변환
	// =========================================================
	fmt.Println("\n=== strconv 패키지 ===")

	// 문자열 → 정수
	num, err := strconv.Atoi("42")
	if err == nil {
		fmt.Printf("Atoi('42') = %d (타입: %T)\n", num, num)
	}

	// 잘못된 변환
	_, err = strconv.Atoi("abc")
	fmt.Println("Atoi('abc') 에러:", err)

	// 정수 → 문자열
	str := strconv.Itoa(123)
	fmt.Printf("Itoa(123) = '%s' (타입: %T)\n", str, str)

	// 실수 변환
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Printf("ParseFloat('3.14159') = %f\n", f)

	fStr := strconv.FormatFloat(f, 'f', 2, 64)
	fmt.Printf("FormatFloat(3.14159, 소수 2자리) = '%s'\n", fStr)

	// bool 변환
	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool('true') = %v\n", b)
}
