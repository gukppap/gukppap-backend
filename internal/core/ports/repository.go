package ports

import (
	"context"

	"github.com/guckppap/gukppap-backend/internal/core/domain"
)

// ShortcutRepository 인터페이스는 Shortcut의 저장소(repository)를 CRUD를 하기 위한 규약입니다.
// ShortcutRepository는 Driven Actor(본 프로젝트에서는 Mysql이 해당됩니다.)가 구현합니다.
type ShorcutRepository interface {
	GetByUrl(ctx context.Context, url string) (*domain.Shortcut, error)
	GetByShortcut(ctx context.Context, shortuct string) (*domain.Shortcut, error)
	Save(context.Context, *domain.Shortcut) error
	Update(context.Context, *domain.Shortcut) error
	Delete(ctx context.Context, url string) error
}
