package main

import "fmt"

// producer: 채널에 값을 보내는 고루틴
// chan<- int: 전송만 가능한 채널 (방향성 채널)
func producer(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i // 채널에 값 전송
		fmt.Printf("전송: %d\n", i)
	}
	close(ch) // 전송 완료 후 채널 닫기 (중요!)
}

// consumer: 채널에서 값을 받는 고루틴
// <-chan int: 수신만 가능한 채널 (방향성 채널)
func consumer(ch <-chan int, results chan<- int) {
	sum := 0
	for val := range ch { // close될 때까지 계속 수신
		sum += val
	}
	results <- sum
}

func main() {
	// =========================================================
	// 언버퍼드 채널 (Unbuffered Channel)
	// 전송자와 수신자가 동시에 준비되어야 전송됩니다.
	// =========================================================
	fmt.Println("=== 언버퍼드 채널 ===")
	ch := make(chan string) // 버퍼 없음

	go func() {
		ch <- "Hello from goroutine!" // 수신자 준비될 때까지 블로킹
	}()

	msg := <-ch // 전송자 준비될 때까지 블로킹
	fmt.Println("받은 메시지:", msg)

	// =========================================================
	// 버퍼드 채널 (Buffered Channel)
	// 버퍼가 꽉 차기 전까지 블로킹 없이 전송 가능
	// =========================================================
	fmt.Println("\n=== 버퍼드 채널 ===")
	buffered := make(chan int, 3) // 버퍼 크기 3

	// 수신자 없어도 버퍼에 쌓임
	buffered <- 10
	buffered <- 20
	buffered <- 30
	fmt.Printf("버퍼 사용량: %d/%d\n", len(buffered), cap(buffered))

	fmt.Println(<-buffered) // 10
	fmt.Println(<-buffered) // 20
	fmt.Println(<-buffered) // 30

	// =========================================================
	// 방향성 채널 + range로 수신
	// =========================================================
	fmt.Println("\n=== 방향성 채널 + range ===")
	dataCh := make(chan int, 5)
	resultCh := make(chan int, 1)

	go producer(dataCh, 5)
	go consumer(dataCh, resultCh)

	total := <-resultCh
	fmt.Printf("0+1+2+3+4 = %d\n", total)

	// =========================================================
	// 채널로 완료 신호 보내기 (done channel 패턴)
	// =========================================================
	fmt.Println("\n=== done channel 패턴 ===")
	done := make(chan struct{}) // 빈 구조체: 신호만 보낼 때 사용 (메모리 0)

	go func() {
		fmt.Println("작업 시작...")
		// 작업 수행
		fmt.Println("작업 완료!")
		close(done) // 완료 신호
	}()

	<-done // 완료 신호 대기
	fmt.Println("메인: 고루틴 작업 확인 완료")
}
