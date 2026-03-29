package homework

import (
	"testing"
)

func TestDeclareVariables(t *testing.T) {
	name, age, height, weight := DeclareVariables()
	if name != "Gopher" {
		t.Errorf("Q1) name: got %q, want %q", name, "Gopher")
	}
	if age != 5 {
		t.Errorf("Q1) age: got %d, want %d", age, 5)
	}
	if height != 171.5 {
		t.Errorf("Q1) height: got %v, want %v", height, 171.5)
	}
	if weight != 68.0 {
		t.Errorf("Q1) weight: got %v, want %v", weight, 68.0)
	}
}

func TestWeekday(t *testing.T) {
	cases := []struct {
		day      Weekday
		weekend  bool
	}{
		{Mon, false},
		{Tue, false},
		{Wed, false},
		{Thu, false},
		{Fri, false},
		{Sat, true},
		{Sun, true},
	}
	for _, c := range cases {
		got := IsWeekend(c.day)
		if got != c.weekend {
			t.Errorf("Q2) IsWeekend(%v): got %v, want %v", c.day, got, c.weekend)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	cases := []struct {
		c, f float64
	}{
		{0, 32},
		{100, 212},
		{37, 98.6},
	}
	for _, c := range cases {
		got := CelsiusToFahrenheit(c.c)
		if got < c.f-0.01 || got > c.f+0.01 {
			t.Errorf("Q3) CelsiusToFahrenheit(%v): got %v, want %v", c.c, got, c.f)
		}
	}
}

func TestFormatProfile(t *testing.T) {
	got := FormatProfile("Gopher", 5, 98.5)
	want := "이름: Gopher, 나이: 5, 점수: 98.50"
	if got != want {
		t.Errorf("Q4) FormatProfile: got %q, want %q", got, want)
	}
}
