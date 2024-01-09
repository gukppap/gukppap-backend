package main

import (
	"github.com/joho/godotenv"
)

func init() {
	// docker-compose와 conatiner관련으로 의해 godotenv.Load()햠수를 여러가지로 받음
	if godotenv.Load("../../.env") != nil {
		if err := godotenv.Load("./.env"); err != nil {
			panic(err)
		}
	}
}

func main() {

}

/*
TODO:
- [V] api 이름을 좀더 간결하고 명확하게 변경


- [V] handler와 repository에 domain을 사용하는 것을 고려해봐야 함 (domain을 사용하니, 어떤 값이 진짜고 가짜인지 잘 모르겠음)


- [V] API 명세 구체적으로 하기
- logging 처리
- helmet 처리
- panic recover 달기
- [V]url에 유효기간을 추가해야 하며, [X] 시간을 더 늘리거나 짧게 하는 url개발
- error handling 개발
- 테스트 코드 작성
- [V] 데이터베이스, 해당 테이블이 존재하는지 확인하는 과정이 필요

- 아 그리고 코드에 주석좀 달아라 제발

- Swaager 달아보기
*/
