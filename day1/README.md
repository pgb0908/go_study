# Day 1: Go 시작하기 - 변수, 타입, 입출력

> 학습 예상 시간: 3~4시간

## 오늘의 학습 목표

오늘 학습을 마치면 다음을 할 수 있습니다:

- [ ] VSCode에서 Go 파일을 작성하고 실행할 수 있다
- [ ] 변수를 3가지 방법으로 선언할 수 있다
- [ ] Go의 기본 타입 6가지를 설명할 수 있다
- [ ] fmt 패키지로 출력과 입력을 처리할 수 있다
- [ ] `const`와 `iota`를 사용해 상수를 정의할 수 있다

## 실행 방법

프로젝트 루트(`go_study/`)에서 실행하세요:

```bash
go run ./day1/01_hello/
go run ./day1/02_variables/
go run ./day1/03_input/
go run ./day1/04_constants/
```

## 핵심 개념 정리

### 1. 패키지와 import

```go
package main   // 실행 가능한 프로그램은 반드시 package main

import "fmt"   // 표준 라이브러리 가져오기
```

### 2. 변수 선언 3가지 방법

```go
// 방법 1: var + 타입 명시
var name string = "홍길동"

// 방법 2: var + 타입 추론 (컴파일러가 타입 결정)
var age = 25

// 방법 3: 짧은 선언 := (함수 안에서만 사용 가능)
city := "서울"
```

### 3. Go의 제로값 (Zero Value)

변수를 선언만 하고 값을 주지 않으면 자동으로 기본값이 됩니다:

| 타입 | 제로값 |
|------|--------|
| int, float64 | 0, 0.0 |
| string | "" (빈 문자열) |
| bool | false |

### 4. 자주 쓰는 fmt 함수

```go
fmt.Println("줄바꿈 포함 출력")
fmt.Printf("이름: %s, 나이: %d\n", name, age)  // 형식 지정 출력
fmt.Scan(&name)   // 입력 받기
```

## 예제별 설명

### 01_hello - Hello World 해부

`package main`이 왜 필요한지, `import`의 역할, `func main()`이 진입점인 이유를 코드 주석으로 설명합니다.

### 02_variables - 변수와 타입

변수 선언 3가지 방법과 제로값 개념을 직접 출력해서 확인합니다.

### 03_input - 사용자 입력

`fmt.Scan`으로 이름을 입력받아 인사하는 프로그램입니다.

### 04_constants - 상수와 iota

`const`로 상수를 정의하고, `iota`로 요일/월을 열거형처럼 만드는 패턴을 배웁니다.

## 실습 과제 (스스로 해보기)

1. **[기본]** 자신의 이름, 나이, 키를 변수에 저장하고 출력해보세요
2. **[응용]** `iota`를 사용해 월(January=1 ~ December=12)을 상수로 정의해보세요
3. **[도전]** 이름과 나이를 입력받아 "안녕하세요, N살 홍길동님!" 형식으로 출력해보세요

## 주의사항 / 흔한 실수

- Go는 선언했지만 **사용하지 않은 변수**가 있으면 컴파일 에러가 납니다
- `:=`는 **함수 안에서만** 사용 가능합니다 (패키지 레벨에서는 `var` 사용)
- `fmt.Scan`으로 입력받을 때는 변수 앞에 `&`(주소 연산자)를 붙여야 합니다

## 참고 자료

- [A Tour of Go - Basics](https://go.dev/tour/basics)
- [Go by Example - Variables](https://gobyexample.com/variables)
- [Go by Example - Constants](https://gobyexample.com/constants)

## 다음 날 예고

Day 2에서는 **제어 흐름**을 배웁니다. Go는 `while`이 없고 `for`만 존재한다는 사실, 알고 계셨나요?
