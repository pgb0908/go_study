# Day 2: 제어 흐름 - if, for, switch

> 학습 예상 시간: 3~4시간

## 오늘의 학습 목표

오늘 학습을 마치면 다음을 할 수 있습니다:

- [ ] Go의 `if` 초기화 구문을 사용할 수 있다
- [ ] `for` 하나로 while처럼, range처럼 사용할 수 있다
- [ ] `switch`에서 `break`가 필요 없는 이유를 설명할 수 있다
- [ ] FizzBuzz를 Go로 작성할 수 있다

## 실행 방법

```bash
go run ./day2/01_if_else/
go run ./day2/02_for_loop/
go run ./day2/03_switch/
go run ./day2/04_fizzbuzz/
```

## 핵심 개념 정리

### 1. if - 초기화 구문

Go의 `if`는 조건 앞에 **짧은 초기화 구문**을 넣을 수 있습니다:

```go
// 일반 if
if x > 0 {
    fmt.Println("양수")
}

// 초기화 구문 포함 (err의 범위가 if 블록 안으로 제한됨)
if err := someFunc(); err != nil {
    fmt.Println("에러:", err)
}
```

### 2. for - Go의 유일한 반복문

Go에는 `while`이 없습니다. `for` 하나로 모든 반복을 처리합니다:

```go
// C 스타일
for i := 0; i < 5; i++ { }

// while 스타일 (조건만)
for x < 100 { }

// 무한 루프
for { }

// range 스타일 (슬라이스/맵 순회)
for i, v := range slice { }
```

### 3. switch - break 없이도 자동 종료

```go
switch day {
case "월요일":
    fmt.Println("주의 시작!")
case "금요일":
    fmt.Println("곧 주말!")
default:
    fmt.Println("평일")
}
// break 없어도 됩니다! Go의 switch는 각 case가 자동으로 끝납니다.
```

## 예제별 설명

### 01_if_else
조건문 기본 문법과 Go만의 초기화 구문 패턴을 배웁니다. 에러 처리에서 자주 쓰이는 패턴입니다.

### 02_for_loop
세 가지 for 사용법 (C스타일, while스타일, range)을 모두 연습합니다.

### 03_switch
switch의 기본 사용법과 표현식 없는 switch (if-else chain 대체)를 배웁니다.

### 04_fizzbuzz
1부터 100까지 3의 배수는 "Fizz", 5의 배수는 "Buzz", 둘 다면 "FizzBuzz"를 출력합니다. 조건문과 반복문을 통합 연습합니다.

## 실습 과제

1. **[기본]** 1부터 50까지의 합을 `for`로 계산해보세요
2. **[응용]** 구구단 2단부터 9단까지 출력해보세요 (중첩 for 사용)
3. **[도전]** 별(`*`)로 5행짜리 삼각형을 출력해보세요

## 주의사항 / 흔한 실수

- Go의 `if`, `for`, `switch` 조건식에는 **괄호가 필요 없습니다** (`if (x > 0)` ← 이렇게 쓰면 Go답지 않음)
- `switch`에서 여러 case를 이어가려면 `fallthrough` 키워드를 명시해야 합니다
- `for range`에서 인덱스만 필요하면 `for i := range slice`, 값만 필요하면 `for _, v := range slice`

## 참고 자료

- [Go by Example - If/Else](https://gobyexample.com/if-else)
- [Go by Example - For](https://gobyexample.com/for)
- [Go by Example - Switch](https://gobyexample.com/switch)

## 다음 날 예고

Day 3에서는 **함수**를 배웁니다. Go 함수는 값을 여러 개 반환할 수 있습니다!
