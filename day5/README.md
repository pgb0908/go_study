# Day 5: 메서드, 인터페이스, 에러 처리

> 학습 예상 시간: 4~5시간

## 오늘의 학습 목표

오늘 학습을 마치면 다음을 할 수 있습니다:

- [ ] 구조체에 메서드를 붙일 수 있다
- [ ] 값 리시버와 포인터 리시버의 차이를 설명할 수 있다
- [ ] 인터페이스를 정의하고 암묵적으로 구현할 수 있다
- [ ] 커스텀 에러 타입을 만들고 처리할 수 있다
- [ ] `fmt.Stringer` 인터페이스를 구현할 수 있다

## 실행 방법

```bash
go run ./day5/01_methods/
go run ./day5/02_interfaces/
go run ./day5/03_error_handling/
go run ./day5/04_stringer/
```

## 핵심 개념 정리

### 1. 메서드 (Method)

Go에는 클래스가 없지만, 구조체에 메서드를 붙일 수 있습니다:

```go
type Circle struct {
    Radius float64
}

// (c Circle): "값 리시버" - Circle의 복사본으로 동작
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// (c *Circle): "포인터 리시버" - 원본을 직접 수정
func (c *Circle) Scale(factor float64) {
    c.Radius *= factor
}
```

### 2. 인터페이스 (Interface)

Go의 인터페이스는 **암묵적**으로 구현됩니다. `implements` 키워드가 없습니다:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Circle이 Area()와 Perimeter()를 구현하면 자동으로 Shape가 됩니다
```

### 3. 에러 처리

```go
// 커스텀 에러 타입
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
```

## 예제별 설명

### 01_methods - 메서드
도형 구조체에 넓이, 둘레 메서드를 붙입니다. 값 리시버와 포인터 리시버의 차이를 실험합니다.

### 02_interfaces - 인터페이스
Shape 인터페이스를 Circle, Rectangle이 암묵적으로 구현하는 다형성 예제입니다.

### 03_error_handling - 에러 처리
`errors.New`, `fmt.Errorf`, 커스텀 에러 타입, `errors.Is`/`errors.As`를 배웁니다.

### 04_stringer - fmt.Stringer
`String() string` 메서드를 구현해서 `fmt.Println`이 구조체를 어떻게 출력하는지 제어합니다.

## 실습 과제

1. **[기본]** `Rectangle` 구조체에 넓이와 둘레를 반환하는 메서드를 만들어보세요
2. **[응용]** `Animal` 인터페이스(`Speak() string`)를 만들고 `Dog`, `Cat`이 구현하게 해보세요
3. **[도전]** 나이가 음수이면 에러를 반환하는 `NewPerson(name string, age int)` 생성자 함수를 만들어보세요

## 주의사항 / 흔한 실수

- **값 리시버**: 메서드 안에서 구조체를 수정해도 원본에 영향 없음
- **포인터 리시버**: 메서드 안에서 구조체를 수정하면 원본도 변경됨
- 구조체의 메서드 중 하나라도 포인터 리시버라면, 나머지도 포인터 리시버로 통일하는 것이 관례
- 인터페이스를 `nil` 체크할 때는 타입도 같이 확인해야 합니다 (인터페이스의 nil은 복잡합니다)

## 참고 자료

- [Go by Example - Methods](https://gobyexample.com/methods)
- [Go by Example - Interfaces](https://gobyexample.com/interfaces)
- [Go by Example - Errors](https://gobyexample.com/errors)

## 다음 날 예고

Day 6에서는 **동시성(Concurrency)**을 배웁니다. Go의 가장 강력한 특성인 고루틴과 채널입니다!
