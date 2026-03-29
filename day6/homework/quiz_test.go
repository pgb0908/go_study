package homework

import (
	"testing"
	"time"
)

func TestParallelSum(t *testing.T) {
	cases := []struct {
		slices [][]int
		want   int
	}{
		{[][]int{{1, 2, 3}, {4, 5}, {6}}, 21},
		{[][]int{{10}, {20}, {30}}, 60},
		{[][]int{{1, 2, 3, 4, 5}}, 15},
	}
	for _, c := range cases {
		got := ParallelSum(c.slices)
		if got != c.want {
			t.Errorf("Q1) ParallelSum(%v): got %d, want %d", c.slices, got, c.want)
		}
	}
}

func TestPipeline(t *testing.T) {
	ch := Generate(2, 3, 4)
	out := Square(ch)

	wants := []int{4, 9, 16}
	for _, want := range wants {
		select {
		case got, ok := <-out:
			if !ok {
				t.Fatalf("Q2) 채널이 너무 일찍 닫혔습니다 (want %d)", want)
			}
			if got != want {
				t.Errorf("Q2) Square: got %d, want %d", got, want)
			}
		case <-time.After(time.Second):
			t.Fatalf("Q2) 타임아웃: 채널에서 값을 받지 못했습니다 (want %d)", want)
		}
	}
	// 채널이 닫혔는지 확인
	select {
	case _, ok := <-out:
		if ok {
			t.Error("Q2) 모든 값 전송 후 채널이 닫혀야 합니다")
		}
	case <-time.After(time.Second):
		t.Error("Q2) 타임아웃: 채널이 닫히지 않았습니다")
	}
}

func TestWithTimeout(t *testing.T) {
	// 값이 있는 경우
	ch := make(chan int, 1)
	ch <- 42
	val, ok := WithTimeout(ch, 100)
	if !ok || val != 42 {
		t.Errorf("Q3) WithTimeout (값 있음): got (%d,%v), want (42,true)", val, ok)
	}

	// 타임아웃 발생
	empty := make(chan int)
	val, ok = WithTimeout(empty, 50)
	if ok || val != 0 {
		t.Errorf("Q3) WithTimeout (타임아웃): got (%d,%v), want (0,false)", val, ok)
	}
}
