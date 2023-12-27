package repository

import (
	"context"
	"gukppap-backend/internal/adapter/repository/mysql/ent/url"
	"gukppap-backend/internal/core/domain"
)

type urlRepository struct {
	db *DB
}

func NewURLRepository(db *DB) *urlRepository {
	return &urlRepository{db}
}

func (ur *urlRepository) CreateURL(ctx context.Context, URL *domain.URL) (*domain.URL, error) {
	_, err := ur.db.Url.Create().
		SetOriginURL(URL.OriginalURL).
		SetShortcutURL(URL.ShortcutURL).
		Save(ctx)
	return URL, err
}

func (ur *urlRepository) GetURLByShortcutURL(ctx context.Context, shortcutURL string) (*domain.URL, error) {
	URL, err := ur.db.Client.Url.Query().Where(
		url.ShortcutURL(shortcutURL),
	).Only(ctx)

	return &domain.URL{ShortcutURL: URL.ShortcutURL, OriginalURL: URL.OriginURL}, err
}

func (ur *urlRepository) GetURLByOrignalURL(ctx context.Context, originlURL string) (*domain.URL, error) {
	URL, err := ur.db.Client.Url.Query().Where(
		url.OriginURL(originlURL),
	).Only(ctx)

	return &domain.URL{ShortcutURL: URL.ShortcutURL, OriginalURL: URL.OriginURL}, err
}

func (ur *urlRepository) UpdateURL(ctx context.Context, shortuctURL string, URL *domain.URL) (*domain.URL, error) {
	_, err := ur.db.Url.Update().Where(
		url.ShortcutURL(URL.ShortcutURL),
	).
		SetOriginURL(URL.OriginalURL).
		SetShortcutURL(URL.ShortcutURL).
		Save(ctx)

	return URL, err
}

func (ur *urlRepository) DeleteByOriginalURL(ctx context.Context, originURL string) error {
	_, err := ur.db.Url.Delete().Where(
		url.OriginURL(originURL),
	).Exec(ctx)

	return err
}

// var name type

// var name = value

// name := value
