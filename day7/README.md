# Day 7: 테스트, 표준 라이브러리, 미니 프로젝트

> 학습 예상 시간: 4~5시간

## 오늘의 학습 목표

오늘 학습을 마치면 다음을 할 수 있습니다:

- [ ] `*_test.go` 파일에 테스트를 작성하고 실행할 수 있다
- [ ] 테이블 주도 테스트 패턴을 사용할 수 있다
- [ ] `strings`, `strconv` 패키지의 주요 함수를 사용할 수 있다
- [ ] `os`, `bufio`로 파일을 읽고 쓸 수 있다
- [ ] 지금까지 배운 내용을 통합해서 CLI 프로그램을 만들 수 있다

## 실행 방법

```bash
# 일반 실행
go run ./day7/02_stdlib_strings/
go run ./day7/03_stdlib_io/
go run ./day7/04_mini_project/

# 테스트 실행 (go run이 아닌 go test!)
go test ./day7/01_testing/
go test -v ./day7/01_testing/        # 상세 출력
go test -run TestAdd ./day7/01_testing/  # 특정 테스트만
```

## 핵심 개념 정리

### 1. 테스트 파일 규칙

- 파일명은 반드시 `_test.go`로 끝나야 합니다
- 테스트 함수는 `Test`로 시작해야 합니다
- 매개변수는 `*testing.T`이어야 합니다

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2,3) = %d; want 5", result)
    }
}
```

### 2. 테이블 주도 테스트 (Table-Driven Tests)

Go에서 가장 권장하는 테스트 패턴입니다:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"양수", 2, 3, 5},
        {"음수", -1, 1, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.expected {
                t.Errorf("got %d, want %d", got, tt.expected)
            }
        })
    }
}
```

### 3. 자주 쓰는 표준 라이브러리

```go
// strings 패키지
strings.Contains("hello", "ell") // true
strings.Split("a,b,c", ",")      // ["a","b","c"]
strings.Join([]string{"a","b"}, "-") // "a-b"
strings.ToUpper("hello")         // "HELLO"
strings.TrimSpace("  hi  ")      // "hi"

// strconv 패키지
strconv.Atoi("42")     // int 42
strconv.Itoa(42)       // string "42"
strconv.ParseFloat("3.14", 64)
```

## 예제별 설명

### 01_testing - 테스트 작성
Add, Divide 함수를 테스트합니다. 테이블 주도 테스트와 에러 케이스 테스트를 포함합니다.

### 02_stdlib_strings - strings/strconv
문자열 처리에 자주 쓰이는 함수들을 실습합니다.

### 03_stdlib_io - os/bufio
파일 읽기/쓰기, 커맨드라인 인자(`os.Args`) 처리를 배웁니다.

### 04_mini_project - 단어 빈도 분석기 (CLI)
텍스트를 입력받아 단어별 등장 횟수를 분석하고 정렬 출력합니다.
Day 1~7의 모든 개념(변수, 제어흐름, 함수, 맵, 구조체, 에러처리)을 종합합니다.

## 실습 과제

1. **[기본]** `Multiply(a, b int) int` 함수와 테스트를 함께 작성해보세요
2. **[응용]** 파일에서 텍스트를 읽어 줄 수를 세는 프로그램을 만들어보세요
3. **[도전]** 미니 프로젝트를 확장해서 가장 빈번한 상위 N개 단어만 출력해보세요

## 주의사항 / 흔한 실수

- 테스트 실행은 `go run`이 아닌 `go test` 입니다
- `t.Error`는 테스트를 계속 진행, `t.Fatal`은 즉시 종료합니다
- 파일 경로는 프로그램 실행 위치 기준입니다 (절대 경로 권장)
- `os.Args[0]`은 프로그램 이름이므로 실제 인자는 `os.Args[1]`부터 시작합니다

## 참고 자료

- [Go by Example - Testing](https://gobyexample.com/testing)
- [Go 공식 testing 패키지](https://pkg.go.dev/testing)
- [Go 공식 strings 패키지](https://pkg.go.dev/strings)

## 수고하셨습니다!

7일 동안 Go의 핵심 개념을 모두 배웠습니다:

| Day | 완료 |
|-----|------|
| 1 - 변수, 타입, 입출력 | ✅ |
| 2 - 제어 흐름 | ✅ |
| 3 - 함수 | ✅ |
| 4 - 자료구조 | ✅ |
| 5 - 메서드, 인터페이스 | ✅ |
| 6 - 동시성 | ✅ |
| 7 - 테스트, 표준 라이브러리 | ✅ |

다음 단계로 추천하는 학습:
- [Go Tour](https://go.dev/tour) 완주
- [Effective Go](https://go.dev/doc/effective_go) 읽기
- 작은 HTTP API 서버 만들기 (`net/http` 패키지)
- Go 오픈소스 프로젝트 코드 읽기
