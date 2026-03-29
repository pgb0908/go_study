# Day 6 Homework

> `quiz.go` 파일의 `TODO` 부분을 완성하세요.
> 완성 후 아래 명령으로 채점합니다:
>
> ```bash
> go test ./day6/homework/ -v
> ```

---

## Q1. 고루틴 + WaitGroup

`ParallelSum(slices [][]int) int` 함수를 완성하세요.

여러 슬라이스를 **고루틴으로 병렬**로 합산하여 전체 합계를 반환합니다.

- `sync.WaitGroup`과 `sync.Mutex` (또는 channel)을 사용하세요
- 각 슬라이스를 별도 고루틴에서 처리해야 합니다

예: `ParallelSum([][]int{{1,2,3},{4,5},{6}})` → `21`

---

## Q2. 채널 - 파이프라인

두 함수를 완성하세요.

`Generate(nums ...int) <-chan int`
- 정수들을 채널에 보내고 채널을 반환합니다

`Square(in <-chan int) <-chan int`
- 입력 채널에서 값을 받아 제곱한 값을 출력 채널에 보냅니다

```go
// 사용 예:
ch := Generate(2, 3, 4)
out := Square(ch)
// out에서 4, 9, 16 순서로 읽힘
```

---

## Q3. select - 타임아웃

`WithTimeout(ch <-chan int, timeoutMs int) (int, bool)` 함수를 완성하세요.

- 채널에서 값을 받으면 `(값, true)` 반환
- `timeoutMs` 밀리초 안에 값이 없으면 `(0, false)` 반환
- `select`와 `time.After`를 사용하세요
