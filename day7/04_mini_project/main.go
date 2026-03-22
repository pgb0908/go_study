// 단어 빈도 분석기 (Word Frequency Analyzer)
// Day 1~7에서 배운 개념 종합:
//   - 변수, 타입 (Day1)
//   - 제어 흐름: for, if (Day2)
//   - 함수, 다중 반환값 (Day3)
//   - 슬라이스, 맵, 구조체 (Day4)
//   - 메서드, 에러 처리 (Day5)
//   - 표준 라이브러리: bufio, os, strings, sort (Day7)
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// WordCount: 단어와 그 빈도를 담는 구조체 (Day4: 구조체)
type WordCount struct {
	Word  string
	Count int
}

// String: fmt.Stringer 구현 (Day5: 인터페이스)
func (wc WordCount) String() string {
	return fmt.Sprintf("%-15s %d회", wc.Word, wc.Count)
}

// analyze: 텍스트를 분석해서 단어 빈도 맵을 반환 (Day3: 함수)
func analyze(text string) map[string]int {
	freq := make(map[string]int) // Day4: 맵

	// 줄 단위로 처리
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords) // 단어 단위로 분리

	for scanner.Scan() {
		word := scanner.Text()
		// 소문자 변환 + 특수문자 제거
		word = strings.ToLower(word)
		word = strings.Trim(word, ".,!?;:\"'()[]")
		if word == "" {
			continue
		}
		freq[word]++ // Day4: 맵 사용
	}
	return freq
}

// topN: 빈도 순으로 정렬된 상위 N개를 반환 (Day4: 슬라이스, sort)
func topN(freq map[string]int, n int) []WordCount {
	// 맵을 슬라이스로 변환
	counts := make([]WordCount, 0, len(freq))
	for word, count := range freq {
		counts = append(counts, WordCount{Word: word, Count: count})
	}

	// 빈도 내림차순 정렬 (같으면 알파벳 순)
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].Count != counts[j].Count {
			return counts[i].Count > counts[j].Count
		}
		return counts[i].Word < counts[j].Word
	})

	// 상위 N개 반환 (Day2: if)
	if n > len(counts) {
		n = len(counts)
	}
	return counts[:n]
}

// readFromStdin: 표준 입력에서 텍스트를 읽기 (Day7: bufio, os)
func readFromStdin() (string, error) {
	var sb strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("분석할 텍스트를 입력하세요.")
	fmt.Println("(입력 종료: 빈 줄 2번 연속 입력)")
	fmt.Println(strings.Repeat("-", 40))

	emptyCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			emptyCount++
			if emptyCount >= 2 {
				break
			}
		} else {
			emptyCount = 0
		}
		sb.WriteString(line)
		sb.WriteString(" ")
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("입력 읽기 실패: %w", err) // Day5: 에러 처리
	}
	return sb.String(), nil
}

func main() {
	fmt.Println("============================")
	fmt.Println("  단어 빈도 분석기 v1.0")
	fmt.Println("============================")

	// 텍스트 입력 받기
	text, err := readFromStdin()
	if err != nil {
		fmt.Println("에러:", err)
		os.Exit(1)
	}

	if strings.TrimSpace(text) == "" {
		fmt.Println("입력된 텍스트가 없습니다.")
		return
	}

	// 분석
	freq := analyze(text)

	// 전체 통계
	totalWords := 0
	for _, count := range freq {
		totalWords += count
	}

	fmt.Printf("\n총 단어 수: %d개\n", totalWords)
	fmt.Printf("고유 단어 수: %d개\n", len(freq))

	// 상위 10개 출력
	top := topN(freq, 10)
	fmt.Printf("\n상위 %d개 단어:\n", len(top))
	fmt.Println(strings.Repeat("-", 25))
	for i, wc := range top {
		fmt.Printf("%2d. %s\n", i+1, wc)
	}
	fmt.Println(strings.Repeat("-", 25))
}
