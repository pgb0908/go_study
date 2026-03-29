# Go 개발 환경 구성 가이드

> 이 문서는 Ubuntu/Linux 환경에서 Go + VSCode 개발 환경을 구성하는 방법을 정리합니다.

---

## 1. Go 설치

### 최신 버전 확인

[https://go.dev/dl/](https://go.dev/dl/) 에서 최신 stable 버전을 확인합니다.

현재 이 프로젝트는 **Go 1.26.1** 기준으로 구성되어 있습니다.

### 설치 (홈 디렉토리에 설치 - sudo 불필요)

```bash
# 1. 다운로드
cd /tmp
wget https://go.dev/dl/go1.26.1.linux-amd64.tar.gz

# 2. 압축 해제 (홈 디렉토리의 go_sdk 폴더에 설치)
mkdir -p ~/go_sdk
tar -C ~/go_sdk -xzf /tmp/go1.26.1.linux-amd64.tar.gz

# 3. 설치 확인
~/go_sdk/go/bin/go version
```

> **참고:** `/usr/local/go`에 설치하려면 `sudo`가 필요합니다.
> 권한이 없는 환경에서는 홈 디렉토리에 설치하면 됩니다.

### PATH 환경변수 등록

`~/.bashrc` 파일에 아래 내용을 추가합니다:

```bash
# ~/.bashrc 에 추가
export GOROOT=$HOME/go_sdk/go       # Go 설치 경로
export GOPATH=$HOME/go              # Go 작업 경로 (다운받은 패키지 등)
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

```bash
# 변경사항 즉시 적용
source ~/.bashrc

# 확인
go version
# 출력 예: go version go1.26.1 linux/amd64
```

### 주요 디렉토리 설명

| 경로 | 설명 |
|------|------|
| `~/go_sdk/go/` | Go 컴파일러, 표준 라이브러리 설치 위치 (GOROOT) |
| `~/go/bin/` | `go install`로 설치한 툴 실행 파일 위치 (GOPATH/bin) |
| `~/go/pkg/` | 컴파일된 패키지 캐시 |

---

## 2. VSCode 확장 설치

### Go 공식 확장 설치

VSCode에서 `Ctrl+Shift+X` → "Go" 검색 → **golang.go** 설치

또는 터미널에서:

```bash
code --install-extension golang.go
```

### Go 툴체인 설치 (VSCode 내부에서)

Go 확장을 설치하면 VSCode가 필요한 툴을 자동으로 설치하라고 안내합니다.

1. `.go` 파일을 열면 우측 하단에 팝업 표시
2. **"Install All"** 클릭

또는 `Ctrl+Shift+P` → `Go: Install/Update Tools` → 전체 선택 → OK

주요 툴:

| 툴 | 역할 |
|----|------|
| `gopls` | 언어 서버 (자동완성, 정의로 이동, 리팩토링) |
| `dlv` (Delve) | 디버거 (브레이크포인트, 변수 검사) |
| `staticcheck` | 코드 정적 분석 |
| `golangci-lint` | 린터 (코드 스타일 검사) |

### dlv (디버거) 수동 설치

팝업이 뜨지 않거나 설치 실패 시 직접 설치:

```bash
go install -v github.com/go-delve/delve/cmd/dlv@latest
```

---

## 3. 프로젝트 초기화

### Go 모듈 초기화

```bash
# 프로젝트 디렉토리에서 실행
cd ~/Desktop/work/go_study/go_study
go mod init go_study
```

`go.mod` 파일이 생성됩니다:

```
module go_study

go 1.26.1
```

### .gitignore 생성

```gitignore
# Go 빌드 결과물
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out

# 의존성 캐시
vendor/

# IDE
.idea/
*.swp
*.swo
```

---

## 4. VSCode 설정

### .vscode/settings.json

```json
{
  "go.goroot": "${env:HOME}/go_sdk/go",
  "go.gopath": "${env:HOME}/go",
  "go.toolsManagement.autoUpdate": true,
  "[go]": {
    "editor.formatOnSave": true,
    "editor.defaultFormatter": "golang.go"
  },
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package",
  "go.useLanguageServer": true
}
```

| 설정 | 설명 |
|------|------|
| `go.goroot` | Go 설치 경로 |
| `go.gopath` | Go 작업 경로 |
| `editor.formatOnSave` | 저장 시 자동 포맷 |
| `go.lintOnSave` | 저장 시 자동 린트 |

### .vscode/launch.json

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Go",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${fileDirname}"
    }
  ]
}
```

> **핵심:** `"program": "${fileDirname}"`
> 현재 VSCode에서 **열려 있는 파일의 디렉토리**를 실행합니다.
> `day1/02_variables/main.go`가 열린 상태에서 F5를 누르면 해당 예제가 실행됩니다.

---

## 5. 코드 실행 방법

### 터미널에서 실행

```bash
# 단일 파일
go run main.go

# 특정 디렉토리 (하위 패키지)
go run ./day1/01_hello/
go run ./day1/02_variables/

# 빌드 후 실행
go build -o myapp .
./myapp
```

### VSCode에서 실행

| 단축키 | 동작 |
|--------|------|
| `Ctrl+F5` | 디버그 없이 실행 |
| `F5` | 디버그 모드로 실행 (브레이크포인트 사용 가능) |
| `F9` | 브레이크포인트 토글 |
| `F10` | 한 줄씩 실행 (Step Over) |
| `F11` | 함수 안으로 들어가기 (Step Into) |

> `main.go` 파일이 열려 있는 상태에서 실행해야 합니다.

### 테스트 실행

```bash
# 특정 패키지 테스트
go test ./day7/01_testing/

# 상세 출력
go test -v ./day7/01_testing/

# 특정 테스트 함수만
go test -run TestAdd ./day7/01_testing/

# 전체 프로젝트 테스트
go test ./...
```

---

## 6. 이 프로젝트 구조

```
go_study/
├── go.mod              # 모듈 정의
├── main.go             # 루트 Hello World
├── SETUP.md            # 이 파일
├── .gitignore
├── .vscode/
│   ├── settings.json   # VSCode Go 설정
│   └── launch.json     # 실행/디버그 설정
│
├── day1/   변수, 타입, 입출력, 상수
├── day2/   if, for, switch
├── day3/   함수, 클로저, defer
├── day4/   슬라이스, 맵, 구조체, 포인터
├── day5/   메서드, 인터페이스, 에러처리
├── day6/   고루틴, 채널, select, sync
└── day7/   테스트, 표준 라이브러리, 미니 프로젝트
```

각 `dayN/` 폴더의 `README.md`를 먼저 읽고 시작하세요.

---

## 7. 자주 쓰는 go 명령어

```bash
go version          # Go 버전 확인
go env              # Go 환경변수 전체 출력
go env GOROOT       # 특정 환경변수 확인

go run .            # 현재 디렉토리 실행
go build .          # 현재 디렉토리 빌드
go test ./...       # 전체 테스트

go get 패키지경로    # 외부 패키지 추가
go mod tidy         # go.mod 정리 (불필요한 의존성 제거)
go mod download     # 의존성 다운로드

go fmt ./...        # 전체 코드 포맷
go vet ./...        # 코드 오류 검사
```

---

## 8. 참고 자료

| 링크 | 설명 |
|------|------|
| [go.dev/dl](https://go.dev/dl/) | Go 공식 다운로드 |
| [go.dev/tour](https://go.dev/tour/) | A Tour of Go (공식 입문 튜토리얼) |
| [gobyexample.com](https://gobyexample.com/) | Go by Example (예제 중심 학습) |
| [pkg.go.dev/std](https://pkg.go.dev/std) | 표준 라이브러리 문서 |
| [go.dev/doc/effective_go](https://go.dev/doc/effective_go) | Effective Go (관용적 Go 작성법) |
