# Day 4: 자료구조 - 배열, 슬라이스, 맵, 구조체, 포인터

> 학습 예상 시간: 4~5시간

## 오늘의 학습 목표

오늘 학습을 마치면 다음을 할 수 있습니다:

- [ ] 배열과 슬라이스의 차이를 설명할 수 있다
- [ ] `make`, `append`, `copy`로 슬라이스를 다룰 수 있다
- [ ] 맵에서 키의 존재 여부를 확인하는 ok 패턴을 사용할 수 있다
- [ ] 구조체를 정의하고 필드에 접근할 수 있다
- [ ] 포인터의 `&`와 `*` 의미를 설명할 수 있다

## 실행 방법

```bash
go run ./day4/01_arrays_slices/
go run ./day4/02_maps/
go run ./day4/03_structs/
go run ./day4/04_pointers/
```

## 핵심 개념 정리

### 1. 슬라이스 (Slice) - 배열보다 자주 씁니다

```go
// 배열: 크기가 고정
arr := [3]int{1, 2, 3}

// 슬라이스: 크기가 동적
s := []int{1, 2, 3}
s = append(s, 4, 5)    // 요소 추가

// make로 생성: make([]타입, 길이, 용량)
s2 := make([]int, 3, 5) // 길이 3, 용량 5
```

### 2. 맵 (Map) - key-value 저장소

```go
m := map[string]int{"사과": 1000, "바나나": 500}

// 존재 여부 확인 (ok 패턴)
price, ok := m["사과"]
if ok {
    fmt.Println("가격:", price)
}
```

### 3. 구조체 (Struct)

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "홍길동", Age: 25}
fmt.Println(p.Name)
```

### 4. 포인터 (Pointer)

```go
x := 42
p := &x  // p는 x의 주소를 담은 포인터
*p = 100 // p가 가리키는 곳(x)의 값을 변경
fmt.Println(x) // 100
```

## 예제별 설명

### 01_arrays_slices - 배열과 슬라이스
슬라이스의 내부 구조(len/cap)와 슬라이싱 표현식, append, copy를 배웁니다.

### 02_maps - 맵
맵 생성, CRUD 연산, ok 패턴, range 순회를 배웁니다.

### 03_structs - 구조체
구조체 정의, 초기화, 익명 구조체, 구조체를 함수에 전달하는 방법을 배웁니다.

### 04_pointers - 포인터
주소 연산자(`&`), 역참조(`*`), 포인터를 함수에 전달해서 값을 수정하는 방법을 배웁니다.

## 실습 과제

1. **[기본]** 학생의 이름과 점수를 맵에 저장하고, 평균을 계산해보세요
2. **[응용]** `Person` 구조체 슬라이스를 만들어 나이 순으로 정렬해보세요 (`sort.Slice` 활용)
3. **[도전]** 슬라이스에서 중복을 제거하는 함수를 맵을 활용해서 만들어보세요

## 주의사항 / 흔한 실수

- 슬라이스는 배열의 뷰(view)입니다. 슬라이스를 복사하면 같은 배열을 가리킵니다 (`copy()` 사용 권장)
- 초기화되지 않은 맵에 값을 넣으면 런타임 패닉 발생: 반드시 `make(map...)` 또는 리터럴로 초기화
- 포인터가 `nil`인 상태에서 역참조하면 패닉 발생

## 참고 자료

- [Go by Example - Slices](https://gobyexample.com/slices)
- [Go by Example - Maps](https://gobyexample.com/maps)
- [Go by Example - Structs](https://gobyexample.com/structs)
- [Go by Example - Pointers](https://gobyexample.com/pointers)

## 다음 날 예고

Day 5에서는 **메서드, 인터페이스, 에러 처리**를 배웁니다. Go의 OOP 방식은 Java와 다릅니다!
