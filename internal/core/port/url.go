package port

import (
	"context"
	"gukppap-backend/internal/core/domain"
)

type URLService interface {
	CreateURL(ctx context.Context, url *domain.URL) (*domain.URL, error)
	GetUrlByShortcutURL(ctx context.Context, shortcutURL string) (*domain.URL, error)
	GetUrlByOrignalURL(ctx context.Context, originlURL string) (*domain.URL, error)
	UpdateURL(ctx context.Context, url domain.URL) (*domain.URL, error)
	DeleteByShortcutURL(ctx context.Context, shortcutURL string) error
	DeleteByOriginalURL(ctx context.Context, originURL string) error
}

type URLRepository interface {
	CreateURL(ctx context.Context, url *domain.URL) (*domain.URL, error)
	GetUrlByShortcutURL(ctx context.Context, shortcutURL string) (*domain.URL, error)
	GetUrlByOrignalURL(ctx context.Context, originlURL string) (*domain.URL, error)
	UpdateURL(ctx context.Context, url domain.URL) (*domain.URL, error)
	DeleteByShortcutURL(ctx context.Context, shortcutURL string) error
	DeleteByOriginalURL(ctx context.Context, originURL string) error
}
