package ports

import (
	"context"

	"github.com/guckppap/gukppap-backend/internal/core/domain"
)

// ShortcutService 인터페이스는 비지니스 로직에 관환 것들의 규약입니다.
// Handler(Driver Actor)로 부터 사용됩니다.
type ShortcutService interface {
	GetByUrl(ctx context.Context, url string) (*domain.Shortcut, error)
	GetByShortcut(ctx context.Context, shortuct string) (*domain.Shortcut, error)
	Save(ctx context.Context, url string) (*domain.Shortcut, error)
	Update(ctx context.Context, shortcut, want string) error
	Delete(ctx context.Context, url string) error
}
