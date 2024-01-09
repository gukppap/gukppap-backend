package domain

import "time"

// NewShortcut 함수는 Shortcut를 생성하여, 본 인스턴스의 주소값을 반환하는 함수 입니다.
// 만료 시간의 기본값은 30일 이므로 현제 시간을 기준으로 만료시간을 계산하여 해당 인스터스를 반환합니다.
func NewShortcut(url string, shortcut string) *Shortcut {
	s := Shortcut{}

	// initialize
	s.Url = url
	s.Shortcut = shortcut
	s.Exp = time.Now().Add(time.Hour * 24 * 30)

	return &s
}

// Shortcut 구조체는 Url의 단축 정보를 같습니다.
type Shortcut struct {
	Url      string    `json:"url"`
	Shortcut string    `json:"shortcut"`
	Exp      time.Time `json:"exp"`
}

// IsValid 함수는 s의 만료시간(exp)과 현재 시간을 비교하여 유효한지 확인하여 참과 거짓을 반환하는 함수 입니다.
func (s *Shortcut) IsValid() bool {
	// Exp가 현제 시간보다 먼저 인가?
	if s.Exp.After(time.Now()) {
		return true
	}
	return false
}
