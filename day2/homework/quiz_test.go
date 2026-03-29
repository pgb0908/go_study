package homework

import (
	"reflect"
	"testing"
)

func TestSumRange(t *testing.T) {
	cases := []struct{ n, want int }{
		{1, 1}, {5, 15}, {10, 55}, {100, 5050},
	}
	for _, c := range cases {
		got := SumRange(c.n)
		if got != c.want {
			t.Errorf("Q1) SumRange(%d): got %d, want %d", c.n, got, c.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 15}
	for _, n := range primes {
		if !IsPrime(n) {
			t.Errorf("Q2) IsPrime(%d): got false, want true", n)
		}
	}
	for _, n := range notPrimes {
		if IsPrime(n) {
			t.Errorf("Q2) IsPrime(%d): got true, want false", n)
		}
	}
}

func TestGetSeason(t *testing.T) {
	cases := []struct {
		month int
		want  string
	}{
		{1, "겨울"}, {2, "겨울"}, {3, "봄"}, {4, "봄"}, {5, "봄"},
		{6, "여름"}, {7, "여름"}, {8, "여름"}, {9, "가을"}, {10, "가을"},
		{11, "가을"}, {12, "겨울"}, {0, "알 수 없음"}, {13, "알 수 없음"},
	}
	for _, c := range cases {
		got := GetSeason(c.month)
		if got != c.want {
			t.Errorf("Q3) GetSeason(%d): got %q, want %q", c.month, got, c.want)
		}
	}
}

func TestMultiplicationTable(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{
		{2, []int{2, 4, 6, 8, 10, 12, 14, 16, 18}},
		{3, []int{3, 6, 9, 12, 15, 18, 21, 24, 27}},
		{9, []int{9, 18, 27, 36, 45, 54, 63, 72, 81}},
	}
	for _, c := range cases {
		got := MultiplicationTable(c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Q4) MultiplicationTable(%d): got %v, want %v", c.n, got, c.want)
		}
	}
}
