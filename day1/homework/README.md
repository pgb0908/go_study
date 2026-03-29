# Day 1 Homework

> `quiz.go` 파일의 `TODO` 부분을 완성하세요.
> 완성 후 아래 명령으로 채점합니다:
>
> ```bash
> go test ./day1/homework/ -v
> ```

---

## Q1. 변수 선언 3가지 방법

`DeclareVariables()` 함수를 완성하세요.

- `var` 키워드로 `name` 선언 (string, 값: `"Gopher"`)
- `:=` 단축 선언으로 `age` 선언 (int, 값: `5`)
- `var` 블록으로 `height`, `weight` 선언 (float64, 값: `171.5`, `68.0`)

반환값: `name, age, height, weight`

---

## Q2. iota 상수

`Weekday` 타입과 상수를 선언하세요.

```
Mon = 1, Tue = 2, Wed = 3, Thu = 4, Fri = 5, Sat = 6, Sun = 7
```

- `iota`를 활용해야 합니다
- `IsWeekend(d Weekday) bool` 함수를 완성하세요 (Sat, Sun이면 true)

---

## Q3. 타입 변환

`CelsiusToFahrenheit(c float64) float64` 함수를 완성하세요.

공식: `F = C * 9/5 + 32`

---

## Q4. 문자열 포맷

`FormatProfile(name string, age int, score float64) string` 함수를 완성하세요.

반환 형식: `"이름: Gopher, 나이: 5, 점수: 98.50"`

(score는 소수점 2자리 고정)
