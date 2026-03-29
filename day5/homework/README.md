# Day 5 Homework

> `quiz.go` 파일의 `TODO` 부분을 완성하세요.
> 완성 후 아래 명령으로 채점합니다:
>
> ```bash
> go test ./day5/homework/ -v
> ```

---

## Q1. 메서드 - 값/포인터 리시버

`BankAccount` 구조체를 완성하세요.

- 필드: `owner string`, `balance float64`
- `NewBankAccount(owner string) *BankAccount`: 생성자 (balance=0)
- `(a *BankAccount) Deposit(amount float64)`: 잔액 증가 (포인터 리시버)
- `(a *BankAccount) Withdraw(amount float64) error`: 잔액 차감, 잔액 부족 시 에러 반환
- `(a BankAccount) Balance() float64`: 잔액 반환 (값 리시버)

---

## Q2. 인터페이스

`Shape` 인터페이스를 구현하세요.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

- `Circle` 구조체 (필드: `Radius float64`)
- `Square` 구조체 (필드: `Side float64`)

두 구조체 모두 `Shape` 인터페이스를 구현하세요.

`TotalArea(shapes []Shape) float64`: 전체 넓이 합계 반환

---

## Q3. 커스텀 에러

`Divide(a, b float64) (float64, error)` 함수를 완성하세요.

- b가 0이면 커스텀 에러 `ErrDivisionByZero` 반환
- 정상이면 a/b 반환

`ErrDivisionByZero`는 `error` 인터페이스를 구현하는 타입이어야 합니다.
