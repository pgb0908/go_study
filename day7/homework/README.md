# Day 7 Homework

> `quiz.go` 파일의 `TODO` 부분을 완성하세요.
> 완성 후 아래 명령으로 채점합니다:
>
> ```bash
> go test ./day7/homework/ -v
> ```

---

## Q1. 테이블 주도 테스트 작성

`Fibonacci(n int) int` 함수를 완성하세요.

- n번째 피보나치 수를 반환합니다
- F(0)=0, F(1)=1, F(2)=1, F(3)=2, F(4)=3, F(5)=5 ...

그리고 `quiz_test.go`에 테이블 주도 테스트를 직접 추가해보세요.

---

## Q2. strings 표준 라이브러리

`ProcessString(s string) string` 함수를 완성하세요.

아래를 순서대로 처리합니다:
1. 앞뒤 공백 제거 (`strings.TrimSpace`)
2. 소문자로 변환 (`strings.ToLower`)
3. 공백을 `_`로 교체 (`strings.ReplaceAll`)

예: `ProcessString("  Hello World  ")` → `"hello_world"`

---

## Q3. 종합 - 성적 처리기

`GradeReport(scores map[string]int) (string, float64)` 함수를 완성하세요.

- 학생 이름 → 점수 맵을 받아
- 최고 점수 학생 이름과 전체 평균 점수를 반환합니다

예:
```go
scores := map[string]int{"Alice": 90, "Bob": 75, "Carol": 88}
GradeReport(scores) // → ("Alice", 84.33...)
```

빈 맵이면 `("", 0.0)` 반환
