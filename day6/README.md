# Day 6: 동시성 - 고루틴, 채널, select, sync

> 학습 예상 시간: 4~5시간

## 오늘의 학습 목표

오늘 학습을 마치면 다음을 할 수 있습니다:

- [ ] `go` 키워드로 고루틴을 시작할 수 있다
- [ ] 채널로 고루틴 간에 데이터를 안전하게 전달할 수 있다
- [ ] 버퍼드 채널과 언버퍼드 채널의 차이를 설명할 수 있다
- [ ] `select`로 여러 채널을 동시에 대기할 수 있다
- [ ] `sync.WaitGroup`과 `sync.Mutex`를 사용할 수 있다

## 실행 방법

```bash
go run ./day6/01_goroutines/
go run ./day6/02_channels/
go run ./day6/03_select/
go run ./day6/04_sync_waitgroup/
```

## 핵심 개념 정리

### 1. 고루틴 (Goroutine)

Go의 경량 스레드입니다. 수천 개를 동시에 실행해도 됩니다:

```go
go func() {
    // 이 코드는 별도 고루틴에서 실행됨
    fmt.Println("고루틴 실행!")
}()
```

### 2. 채널 (Channel)

고루틴 간 통신 수단입니다. "메모리를 공유해서 통신하지 말고, 통신으로 공유하라":

```go
ch := make(chan int)       // 언버퍼드 채널
ch := make(chan int, 5)    // 버퍼드 채널 (크기 5)

ch <- 42    // 채널에 전송
val := <-ch // 채널에서 수신
```

### 3. select

여러 채널 중 준비된 것을 선택합니다:

```go
select {
case msg := <-ch1:
    fmt.Println("ch1:", msg)
case msg := <-ch2:
    fmt.Println("ch2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("타임아웃!")
}
```

### 4. sync.WaitGroup

여러 고루틴이 모두 끝날 때까지 기다립니다:

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // 작업
}()
wg.Wait() // 모든 고루틴이 Done()을 호출할 때까지 대기
```

## 예제별 설명

### 01_goroutines - 고루틴
`go` 키워드로 고루틴을 시작하고, `time.Sleep`의 문제점을 확인합니다.

### 02_channels - 채널
언버퍼드/버퍼드 채널, 방향성 채널, `range`로 채널 수신을 배웁니다.

### 03_select - select문
여러 채널 동시 대기, 타임아웃 패턴, non-blocking 채널 연산을 배웁니다.

### 04_sync_waitgroup - sync 패키지
WaitGroup으로 고루틴 완료 대기, Mutex로 공유 자원 보호, race condition을 실험합니다.

## 실습 과제

1. **[기본]** 고루틴 3개를 실행해서 각각 다른 메시지를 채널로 보내고, 메인에서 모두 받아보세요
2. **[응용]** 1부터 100까지의 합을 고루틴 2개로 나눠서 계산해보세요 (1~50, 51~100)
3. **[도전]** `select`와 `time.After`를 사용해 3초 타임아웃이 있는 작업을 구현해보세요

## 주의사항 / 흔한 실수

- `time.Sleep`으로 고루틴을 기다리는 것은 위험합니다. `WaitGroup` 또는 채널을 사용하세요
- 언버퍼드 채널에 전송하면 수신자가 없으면 **블로킹**됩니다 (데드락 주의)
- `go test -race ./...` 로 race condition을 감지할 수 있습니다
- 채널을 `close()`한 후에도 이미 있는 값은 계속 읽을 수 있습니다

## 참고 자료

- [Go by Example - Goroutines](https://gobyexample.com/goroutines)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Go by Example - Select](https://gobyexample.com/select)
- [Go by Example - WaitGroups](https://gobyexample.com/waitgroups)

## 다음 날 예고

Day 7에서는 **테스트, 표준 라이브러리, 미니 프로젝트**로 한 주를 마무리합니다!
