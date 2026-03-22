package main

import (
	"fmt"
	"sync"
)

// =========================================================
// sync.WaitGroup: 여러 고루틴이 끝날 때까지 기다리기
// =========================================================
func withWaitGroup() {
	fmt.Println("=== sync.WaitGroup ===")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 카운터 +1 (고루틴 시작 전에 호출!)
		go func(id int) {
			defer wg.Done() // 고루틴이 끝날 때 카운터 -1 (defer로 보장)
			fmt.Printf("고루틴 %d 작업 중...\n", id)
		}(i)
	}

	wg.Wait() // 카운터가 0이 될 때까지 대기
	fmt.Println("모든 고루틴 완료!\n")
}

// =========================================================
// Race Condition (경쟁 조건): 여러 고루틴이 공유 변수를 동시에 수정
// =========================================================
func withoutMutex() {
	fmt.Println("=== Mutex 없이 (Race Condition 발생) ===")
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // 여러 고루틴이 동시에 접근! 결과가 불확실
		}()
	}

	wg.Wait()
	fmt.Printf("예상: 1000, 실제: %d (불확실할 수 있음)\n\n", counter)
}

// =========================================================
// sync.Mutex: 공유 자원을 한 번에 하나의 고루틴만 접근하도록
// =========================================================
func withMutex() {
	fmt.Println("=== sync.Mutex 사용 ===")
	counter := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()   // 잠금: 다른 고루틴은 대기
			counter++   // 안전하게 수정
			mu.Unlock() // 잠금 해제
		}()
	}

	wg.Wait()
	fmt.Printf("예상: 1000, 실제: %d (항상 정확)\n\n", counter)
}

// =========================================================
// sync.RWMutex: 읽기는 여러 고루틴이 동시에, 쓰기는 하나만
// =========================================================
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]int)}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock() // 읽기 잠금 (여러 고루틴 동시 읽기 가능)
	defer sm.mu.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

func main() {
	withWaitGroup()
	withoutMutex()
	withMutex()

	// RWMutex 예제
	fmt.Println("=== sync.RWMutex (안전한 맵) ===")
	sm := NewSafeMap()
	var wg sync.WaitGroup

	// 쓰기 고루틴
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			sm.Set(key, n*10)
		}(i)
	}

	wg.Wait()

	// 읽기
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := sm.Get(key); ok {
			fmt.Printf("%s = %d\n", key, val)
		}
	}
}
