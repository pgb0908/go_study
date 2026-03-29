package homework

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	cases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2, 3, 1, 4}, []int{1, 2, 3, 4}},
		{[]int{5, 5, 5}, []int{5}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		got := Unique(c.nums)
		if len(c.want) == 0 && len(got) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Q1) Unique(%v): got %v, want %v", c.nums, got, c.want)
		}
	}
}

func TestWordCount(t *testing.T) {
	cases := []struct {
		s    string
		want map[string]int
	}{
		{"go is go", map[string]int{"go": 2, "is": 1}},
		{"hello world", map[string]int{"hello": 1, "world": 1}},
		{"a a a", map[string]int{"a": 3}},
	}
	for _, c := range cases {
		got := WordCount(c.s)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Q2) WordCount(%q): got %v, want %v", c.s, got, c.want)
		}
	}
}

func TestRectangle(t *testing.T) {
	r := NewRectangle(4, 5)
	if r.Width != 4 || r.Height != 5 {
		t.Errorf("Q3) NewRectangle(4,5): got {%v,%v}, want {4,5}", r.Width, r.Height)
	}
	if got := Area(r); got != 20 {
		t.Errorf("Q3) Area: got %v, want 20", got)
	}
	if got := Perimeter(r); got != 18 {
		t.Errorf("Q3) Perimeter: got %v, want 18", got)
	}
}

func TestDouble(t *testing.T) {
	x := 5
	Double(&x)
	if x != 10 {
		t.Errorf("Q4) Double(&5): got %d, want 10", x)
	}
	y := -3
	Double(&y)
	if y != -6 {
		t.Errorf("Q4) Double(&-3): got %d, want -6", y)
	}
}
