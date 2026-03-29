package homework

import (
	"reflect"
	"testing"
)

func TestMinMax(t *testing.T) {
	cases := []struct {
		nums     []int
		min, max int
	}{
		{[]int{3, 1, 4, 1, 5, 9}, 1, 9},
		{[]int{-5, 0, 5}, -5, 5},
		{[]int{7}, 7, 7},
	}
	for _, c := range cases {
		gotMin, gotMax := MinMax(c.nums)
		if gotMin != c.min || gotMax != c.max {
			t.Errorf("Q1) MinMax(%v): got (%d,%d), want (%d,%d)", c.nums, gotMin, gotMax, c.min, c.max)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{10, 20}, 30},
		{[]int{}, 0},
		{[]int{-1, 1}, 0},
	}
	for _, c := range cases {
		got := Sum(c.nums...)
		if got != c.want {
			t.Errorf("Q2) Sum(%v): got %d, want %d", c.nums, got, c.want)
		}
	}
}

func TestMakeCounter(t *testing.T) {
	counter := MakeCounter()
	if counter == nil {
		t.Fatal("Q3) MakeCounter() returned nil")
	}
	for i := 1; i <= 5; i++ {
		got := counter()
		if got != i {
			t.Errorf("Q3) counter() call %d: got %d, want %d", i, got, i)
		}
	}
	// 독립적인 카운터인지 확인
	counter2 := MakeCounter()
	got := counter2()
	if got != 1 {
		t.Errorf("Q3) 새 counter() 첫 호출: got %d, want 1", got)
	}
}

func TestDeferOrder(t *testing.T) {
	got := DeferOrder()
	want := []string{"third", "second", "first"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Q4) DeferOrder(): got %v, want %v", got, want)
	}
}
