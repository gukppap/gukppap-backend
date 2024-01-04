package ports

import "github.com/guckppap/gukppap-backend/internal/core/domain"

// ShortcutService 인터페이스는 비지니스 로직에 관환 것들의 규약입니다.
// Handler(Driver Actor)로 부터 사용됩니다.
type ShortcutService interface {
	GetByUrl(url string) (*domain.Shortcut, error)
	GetByShortcut(shortuct string) (*domain.Shortcut, error)
	Save(*domain.Shortcut) error
	Update(*domain.Shortcut) error
	Delete(url string) error
}
