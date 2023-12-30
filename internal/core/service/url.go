package service

import (
	"context"
	"gukppap-backend/internal/core/domain"
	"gukppap-backend/internal/core/port"

	"github.com/google/uuid"
)

// implements port.URLService interface
type URLService struct {
	repo port.URLRepository
}

func NewURLService(repo port.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (us *URLService) CreateURL(ctx context.Context, URL *domain.URL) (*domain.URL, error) {

	uuid := ([]rune(uuid.New().String()))[:7]
	URL.ShortcutURL = string(uuid)

	resURL, err := us.repo.CreateURL(ctx, URL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) GetURLByShortcutURL(ctx context.Context, shortcutURL string) (*domain.URL, error) {
	resURL, err := us.repo.GetURLByShortcutURL(ctx, shortcutURL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) GetURLByOrignalURL(ctx context.Context, originlURL string) (*domain.URL, error) {
	resURL, err := us.repo.GetURLByOrignalURL(ctx, originlURL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) UpdateURL(ctx context.Context, shortcutURL string, URL *domain.URL) (*domain.URL, error) {
	resURL, err := us.repo.UpdateURL(ctx, shortcutURL, URL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) DeleteByOriginalURL(ctx context.Context, originURL string) error {
	ctx.Deadline()
	err := us.repo.DeleteByOriginalURL(ctx, originURL)
	if err != nil {
		return err
	}

	return nil
}
