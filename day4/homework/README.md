# Day 4 Homework

> `quiz.go` 파일의 `TODO` 부분을 완성하세요.
> 완성 후 아래 명령으로 채점합니다:
>
> ```bash
> go test ./day4/homework/ -v
> ```

---

## Q1. 슬라이스 - 중복 제거

`Unique(nums []int) []int` 함수를 완성하세요.

슬라이스에서 중복을 제거하고 **처음 등장한 순서**를 유지해 반환합니다.

예: `Unique([]int{1,2,2,3,1,4})` → `[1, 2, 3, 4]`

---

## Q2. 맵 - 단어 빈도 카운터

`WordCount(s string) map[string]int` 함수를 완성하세요.

문자열을 공백으로 분리해 각 단어의 등장 횟수를 맵으로 반환합니다.

예: `WordCount("go is go")` → `map[go:2 is:1]`

---

## Q3. 구조체

`Rectangle` 구조체를 정의하고 아래 함수를 완성하세요.

- 필드: `Width float64`, `Height float64`
- `NewRectangle(w, h float64) Rectangle`: 생성자
- `Area(r Rectangle) float64`: 넓이 반환 (Width × Height)
- `Perimeter(r Rectangle) float64`: 둘레 반환 (2 × (Width + Height))

---

## Q4. 포인터

`Double(n *int)` 함수를 완성하세요.

포인터로 전달된 값을 2배로 만듭니다 (반환값 없음).

```go
x := 5
Double(&x)
// x == 10
```
