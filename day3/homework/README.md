# Day 3 Homework

> `quiz.go` 파일의 `TODO` 부분을 완성하세요.
> 완성 후 아래 명령으로 채점합니다:
>
> ```bash
> go test ./day3/homework/ -v
> ```

---

## Q1. 다중 반환값

`MinMax(nums []int) (int, int)` 함수를 완성하세요.

슬라이스에서 최솟값과 최댓값을 반환합니다.

예: `MinMax([]int{3,1,4,1,5,9})` → `(1, 9)`

---

## Q2. 가변 인자

`Sum(nums ...int) int` 함수를 완성하세요.

인자 개수에 관계없이 모두 더한 값을 반환합니다.

예: `Sum(1, 2, 3)` → `6`, `Sum(10, 20)` → `30`

---

## Q3. 클로저 - 카운터

`MakeCounter() func() int` 함수를 완성하세요.

호출할 때마다 1씩 증가하는 카운터 함수를 반환합니다.

```go
counter := MakeCounter()
counter() // → 1
counter() // → 2
counter() // → 3
```

---

## Q4. defer 순서

`DeferOrder() []string` 함수를 완성하세요.

`defer`를 사용해서 `"first"`, `"second"`, `"third"` 를 등록하고,
실제 실행 순서대로 슬라이스에 담아 반환합니다.

힌트: defer는 **LIFO** (후입선출) 순서로 실행됩니다.

예상 반환값: `["third", "second", "first"]`
