# 빌드 스테이지 
FROM golang:1.21-alpine3.19 AS builder

# 작업 공간 설정 
WORKDIR /app

# 소스코드를 이미지에 추가 
COPY . .

# 의존성 다운
RUN go mod download

# 실행파일을 위해 빌드
RUN go build -o ./bin/gukppap ./cmd/http/gukppap.go

# 마지막 스테이지 (베이스 이미지)
FROM alpine:3.19 AS final

# 작업 공간 설정
WORKDIR /app

# 실행파일을 복사 
COPY --from=builder /app/bin/gukppap ./
COPY --from=builder /app/.env ./

# 포트 설정 
EXPOSE 8080

# 시작시 실행되는 명령어 
ENTRYPOINT [ "./gukppap" ]

