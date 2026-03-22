# Day 3: 함수 - 다중 반환값, 가변 인자, 클로저, defer

> 학습 예상 시간: 3~4시간

## 오늘의 학습 목표

오늘 학습을 마치면 다음을 할 수 있습니다:

- [ ] Go 함수의 기본 문법을 작성할 수 있다
- [ ] 함수에서 값을 여러 개 반환할 수 있다
- [ ] 가변 인자 함수(`...`)를 정의하고 사용할 수 있다
- [ ] 클로저의 개념을 설명하고 활용할 수 있다
- [ ] `defer`가 언제, 왜 쓰이는지 설명할 수 있다

## 실행 방법

```bash
go run ./day3/01_functions/
go run ./day3/02_multiple_return/
go run ./day3/03_variadic/
go run ./day3/04_closures/
```

## 핵심 개념 정리

### 1. 함수 선언

```go
// func 함수명(매개변수 타입) 반환타입 { ... }
func add(a int, b int) int {
    return a + b
}

// 같은 타입 매개변수는 묶어서 선언 가능
func add(a, b int) int {
    return a + b
}
```

### 2. 다중 반환값 (Go의 핵심 특성!)

```go
// 두 값을 동시에 반환
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("0으로 나눌 수 없습니다")
    }
    return a / b, nil
}

// 호출 시
result, err := divide(10, 3)
if err != nil {
    fmt.Println("에러:", err)
}
```

### 3. 클로저 (Closure)

함수가 자신이 선언된 환경의 변수를 "포획(capture)"해서 계속 참조하는 기능입니다:

```go
func makeCounter() func() int {
    count := 0                 // 이 변수를 클로저가 포획함
    return func() int {
        count++
        return count
    }
}
```

### 4. defer

`defer`가 붙은 함수 호출은 현재 함수가 끝날 때 실행됩니다:

```go
func readFile() {
    f, _ := os.Open("file.txt")
    defer f.Close()   // readFile이 끝날 때 자동으로 실행
    // ... 파일 처리
}
```

## 예제별 설명

### 01_functions - 기본 함수
함수 선언 문법, 재귀 함수(팩토리얼), 함수를 변수에 저장하는 방법을 배웁니다.

### 02_multiple_return - 다중 반환값
Go의 에러 처리 패턴의 기초가 되는 다중 반환값을 배웁니다.

### 03_variadic - 가변 인자
`...` 문법으로 개수가 정해지지 않은 인자를 받는 함수를 만듭니다.

### 04_closures - 클로저와 defer
함수를 값처럼 다루는 방법과, `defer`의 실행 순서를 배웁니다.

## 실습 과제

1. **[기본]** 두 정수를 받아 합, 차, 곱, 몫을 한 번에 반환하는 함수를 만들어보세요
2. **[응용]** 가변 인자로 정수 여러 개를 받아 최솟값과 최댓값을 반환하는 함수를 만들어보세요
3. **[도전]** 클로저를 사용해 "곱하기 N" 함수를 반환하는 `makeMultiplier(n int)` 함수를 만들어보세요

## 주의사항 / 흔한 실수

- 다중 반환값을 받을 때 사용하지 않는 값은 `_`로 버려야 합니다: `result, _ := divide(10, 3)`
- `defer`는 LIFO(후입선출) 순서로 실행됩니다 (여러 개면 마지막에 선언된 게 먼저 실행)
- 가변 인자 함수에 슬라이스를 펼쳐서 전달할 때는 `func(slice...)` 형식을 씁니다

## 참고 자료

- [Go by Example - Functions](https://gobyexample.com/functions)
- [Go by Example - Multiple Return Values](https://gobyexample.com/multiple-return-values)
- [Go by Example - Closures](https://gobyexample.com/closures)
- [Go by Example - Defer](https://gobyexample.com/defer)

## 다음 날 예고

Day 4에서는 **자료구조**를 배웁니다. 슬라이스, 맵, 구조체, 포인터까지!
