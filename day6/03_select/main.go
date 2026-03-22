package main

import (
	"fmt"
	"time"
)

func main() {
	// =========================================================
	// 기본 select: 여러 채널 중 준비된 것을 선택
	// =========================================================
	fmt.Println("=== 기본 select ===")
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "하나"
	ch2 <- "둘"

	// 두 채널 모두 준비됨 - 랜덤으로 선택됩니다
	select {
	case msg := <-ch1:
		fmt.Println("ch1에서 받음:", msg)
	case msg := <-ch2:
		fmt.Println("ch2에서 받음:", msg)
	}

	// =========================================================
	// select + default: Non-blocking 채널 연산
	// =========================================================
	fmt.Println("\n=== Non-blocking select ===")
	ch3 := make(chan int)

	select {
	case val := <-ch3:
		fmt.Println("받음:", val)
	default:
		fmt.Println("채널에 데이터 없음 (non-blocking)") // 바로 실행됨
	}

	// =========================================================
	// 타임아웃 패턴 (실무에서 매우 자주 사용!)
	// =========================================================
	fmt.Println("\n=== 타임아웃 패턴 ===")
	slowCh := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second) // 2초 걸리는 작업
		slowCh <- "느린 응답"
	}()

	select {
	case result := <-slowCh:
		fmt.Println("결과:", result)
	case <-time.After(1 * time.Second): // 1초 타임아웃
		fmt.Println("타임아웃! 1초 초과")
	}

	// =========================================================
	// 취소 패턴: done 채널로 고루틴 종료
	// =========================================================
	fmt.Println("\n=== 취소 패턴 ===")
	done := make(chan struct{})
	work := make(chan int)

	// 일하는 고루틴
	go func() {
		for i := 0; ; i++ {
			select {
			case <-done: // 종료 신호
				fmt.Println("고루틴: 종료 신호 받음")
				return
			case work <- i: // 값 전송
			}
		}
	}()

	// 5개만 받고 취소
	for i := 0; i < 5; i++ {
		fmt.Printf("받음: %d\n", <-work)
	}
	close(done) // 종료 신호
	time.Sleep(10 * time.Millisecond)
	fmt.Println("완료")
}
