package ports

import "github.com/guckppap/gukppap-backend/internal/core/domain"

// ShortcutRepository 인터페이스는 Shortcut의 저장소(repository)를 CRUD를 하기 위한 규약입니다.
// ShortcutRepository는 Driven Actor(본 프로젝트에서는 Mysql이 해당됩니다.)가 구현합니다.
type ShorcutRepository interface {
	GetByUrl(url string) (*domain.Shortcut, error)
	GetByShortcut(shortuct string) (*domain.Shortcut, error)
	Save(*domain.Shortcut) error
	Update(*domain.Shortcut) error
	Delete(url string) error
}
