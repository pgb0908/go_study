package homework

import (
	"errors"
	"math"
)

// Q1. BankAccount 구조체와 메서드를 완성하세요.
type BankAccount struct {
	// TODO: owner string, balance float64
}

func NewBankAccount(owner string) *BankAccount {
	// TODO
	return nil
}

func (a *BankAccount) Deposit(amount float64) {
	// TODO: 포인터 리시버 - balance 증가
}

func (a *BankAccount) Withdraw(amount float64) error {
	// TODO: 잔액 부족이면 errors.New("잔액이 부족합니다") 반환
	_ = errors.New // 이 줄은 삭제하세요
	return nil
}

func (a BankAccount) Balance() float64 {
	// TODO: 값 리시버 - balance 반환
	return 0
}

// Q2. Shape 인터페이스
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle 구조체와 메서드를 완성하세요. (math.Pi 사용)
type Circle struct {
	// TODO: Radius float64
}

func (c Circle) Area() float64 {
	// TODO: π * r²
	_ = math.Pi // 이 줄은 삭제하세요
	return 0
}

func (c Circle) Perimeter() float64 {
	// TODO: 2 * π * r
	return 0
}

// Square 구조체와 메서드를 완성하세요.
type Square struct {
	// TODO: Side float64
}

func (s Square) Area() float64 {
	// TODO: side²
	return 0
}

func (s Square) Perimeter() float64 {
	// TODO: 4 * side
	return 0
}

// TotalArea: Shape 슬라이스의 넓이 합계를 반환합니다.
func TotalArea(shapes []Shape) float64 {
	// TODO
	return 0
}

// Q3. 커스텀 에러 타입을 선언하세요.
// type DivisionError struct { ... }
// func (e DivisionError) Error() string { ... }

var ErrDivisionByZero = errors.New("0으로 나눌 수 없습니다") // 기본 제공, 커스텀 타입으로 바꿔보세요

// Divide: b가 0이면 ErrDivisionByZero 반환, 아니면 a/b 반환
func Divide(a, b float64) (float64, error) {
	// TODO
	return 0, nil
}
