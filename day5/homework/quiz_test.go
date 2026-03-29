package homework

import (
	"math"
	"testing"
)

func TestBankAccount(t *testing.T) {
	acc := NewBankAccount("Gopher")
	if acc == nil {
		t.Fatal("Q1) NewBankAccount returned nil")
	}
	if acc.Balance() != 0 {
		t.Errorf("Q1) 초기 잔액: got %v, want 0", acc.Balance())
	}
	acc.Deposit(1000)
	if acc.Balance() != 1000 {
		t.Errorf("Q1) Deposit 후 잔액: got %v, want 1000", acc.Balance())
	}
	err := acc.Withdraw(300)
	if err != nil {
		t.Errorf("Q1) Withdraw(300) 에러: %v", err)
	}
	if acc.Balance() != 700 {
		t.Errorf("Q1) Withdraw 후 잔액: got %v, want 700", acc.Balance())
	}
	err = acc.Withdraw(999)
	if err == nil {
		t.Error("Q1) 잔액 부족 시 에러가 반환되어야 합니다")
	}
	if acc.Balance() != 700 {
		t.Errorf("Q1) 실패한 Withdraw 후 잔액: got %v, want 700", acc.Balance())
	}
}

func TestShapes(t *testing.T) {
	c := Circle{Radius: 5}
	wantArea := math.Pi * 25
	wantPerim := 2 * math.Pi * 5
	if math.Abs(c.Area()-wantArea) > 0.001 {
		t.Errorf("Q2) Circle.Area(): got %v, want %v", c.Area(), wantArea)
	}
	if math.Abs(c.Perimeter()-wantPerim) > 0.001 {
		t.Errorf("Q2) Circle.Perimeter(): got %v, want %v", c.Perimeter(), wantPerim)
	}

	s := Square{Side: 4}
	if s.Area() != 16 {
		t.Errorf("Q2) Square.Area(): got %v, want 16", s.Area())
	}
	if s.Perimeter() != 16 {
		t.Errorf("Q2) Square.Perimeter(): got %v, want 16", s.Perimeter())
	}

	shapes := []Shape{Square{Side: 3}, Square{Side: 4}}
	total := TotalArea(shapes)
	if total != 25 {
		t.Errorf("Q2) TotalArea: got %v, want 25", total)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Q3) Divide(10,2) 에러: %v", err)
	}
	if result != 5 {
		t.Errorf("Q3) Divide(10,2): got %v, want 5", result)
	}
	_, err = Divide(5, 0)
	if err == nil {
		t.Error("Q3) Divide(5,0): 에러가 반환되어야 합니다")
	}
}
