package port

import (
	"context"
	"gukppap-backend/internal/core/domain"
)

type URLService interface {
	CreateURL(ctx context.Context, URL *domain.URL) (*domain.URL, error)
	GetURLByShortcutURL(ctx context.Context, shortcutURL string) (*domain.URL, error)
	GetURLByOrignalURL(ctx context.Context, originlURL string) (*domain.URL, error)
	UpdateURL(ctx context.Context, URL domain.URL) (*domain.URL, error)
	DeleteByShortcutURL(ctx context.Context, shortcutURL string) error
	DeleteByOriginalURL(ctx context.Context, originURL string) error
}

type URLRepository interface {
	CreateURL(ctx context.Context, URL *domain.URL) (*domain.URL, error)
	GetURLByShortcutURL(ctx context.Context, shortcutURL string) (*domain.URL, error)
	GetURLByOrignalURL(ctx context.Context, originlURL string) (*domain.URL, error)
	UpdateURL(ctx context.Context, URL domain.URL) (*domain.URL, error)
	DeleteByShortcutURL(ctx context.Context, shortcutURL string) error
	DeleteByOriginalURL(ctx context.Context, originURL string) error
}
