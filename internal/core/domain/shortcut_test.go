package domain_test

import (
	"testing"
	"time"

	"github.com/guckppap/gukppap-backend/internal/core/domain"
)

// TestShortcut_ISValid 함수는 Shorcut 구초제의 IsValid함수가 정상적으로 동작하는 확인하는 함수 입니다.
func TestShortcut_IsValid(t *testing.T) {
	// shortcut 구조체 생성
	shortcut := &domain.Shortcut{
		// 만료 시간을 현재 시간 기준으로 24시간 설정
		Exp: time.Now().Add(time.Hour * 24),
	}

	if shortcut.IsValid() == false {
		t.Fatal("참이 나와야 하는데 거짓이 나왔어요. 이것은 만료한데 만료되지 않았다고 하는 것과 같아요.")

	}

}
