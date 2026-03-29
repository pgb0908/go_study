# Day 2 Homework

> `quiz.go` 파일의 `TODO` 부분을 완성하세요.
> 완성 후 아래 명령으로 채점합니다:
>
> ```bash
> go test ./day2/homework/ -v
> ```

---

## Q1. for 루프 - 합산

`SumRange(n int) int` 함수를 완성하세요.

1부터 n까지의 합을 반환합니다.

예: `SumRange(10)` → `55`

---

## Q2. 소수 판별

`IsPrime(n int) bool` 함수를 완성하세요.

- n이 소수이면 true, 아니면 false
- 힌트: 2부터 √n 까지 나누어 떨어지는지 확인

예: `IsPrime(7)` → `true`, `IsPrime(9)` → `false`

---

## Q3. switch - 계절 판별

`GetSeason(month int) string` 함수를 완성하세요.

- 3, 4, 5월 → `"봄"`
- 6, 7, 8월 → `"여름"`
- 9, 10, 11월 → `"가을"`
- 12, 1, 2월 → `"겨울"`
- 그 외 → `"알 수 없음"`

---

## Q4. 중첩 루프 - 구구단

`MultiplicationTable(n int) []int` 함수를 완성하세요.

n단의 결과를 슬라이스로 반환합니다 (1~9 곱한 값).

예: `MultiplicationTable(3)` → `[3, 6, 9, 12, 15, 18, 21, 24, 27]`
