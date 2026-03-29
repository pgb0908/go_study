package homework

import (
	"sync"
	"time"
)

// Q1. 여러 슬라이스를 고루틴으로 병렬 합산합니다.
func ParallelSum(slices [][]int) int {
	// TODO: sync.WaitGroup과 sync.Mutex 사용
	var wg sync.WaitGroup
	var mu sync.Mutex
	total := 0

	for _, s := range slices {
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			// TODO: nums의 합을 구해서 total에 더하세요 (mu.Lock/Unlock 사용)
			_ = mu
		}(s)
	}

	wg.Wait()
	return total
}

// Q2. 정수들을 채널에 보내고 채널을 반환합니다.
func Generate(nums ...int) <-chan int {
	// TODO: 고루틴으로 채널에 nums 전송 후 close
	ch := make(chan int)
	_ = ch
	return ch
}

// Square: 입력 채널 값을 제곱해서 출력 채널로 보냅니다.
func Square(in <-chan int) <-chan int {
	// TODO: 고루틴으로 in 채널에서 읽어 제곱 후 out 채널로 전송
	out := make(chan int)
	_ = out
	return out
}

// Q3. 타임아웃 내에 채널에서 값을 받으면 (값, true), 아니면 (0, false) 반환.
func WithTimeout(ch <-chan int, timeoutMs int) (int, bool) {
	// TODO: select + time.After 사용
	_ = time.After
	return 0, false
}
