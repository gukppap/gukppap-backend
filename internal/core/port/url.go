package port

import (
	"gukppap-backend/internal/core/domain"
)

type URLService interface {
	CreateURL(URL *domain.URL) (*domain.URL, error)
	GetURLByShortcutURL(shortcutURL string) (*domain.URL, error)
	GetURLByOrignalURL(originlURL string) (*domain.URL, error)
	UpdateURL(URL domain.URL) (*domain.URL, error)
	DeleteByShortcutURL(shortcutURL string) error
	DeleteByOriginalURL(originURL string) error
}

type URLRepository interface {
	CreateURL(URL *domain.URL) (*domain.URL, error)
	GetURLByShortcutURL(shortcutURL string) (*domain.URL, error)
	GetURLByOrignalURL(originlURL string) (*domain.URL, error)
	UpdateURL(URL domain.URL) (*domain.URL, error)
	DeleteByShortcutURL(shortcutURL string) error
	DeleteByOriginalURL(originURL string) error
}
