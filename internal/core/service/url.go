package service

import (
	"gukppap-backend/internal/core/domain"
	"gukppap-backend/internal/core/port"

	"github.com/dchest/uniuri"
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

func (us *URLService) CreateURL(URL *domain.URL) (*domain.URL, error) {
	URL.ShortcutURL = uniuri.NewLenChars(6, []byte(URL.OriginalURL))

	resURL, err := us.repo.CreateURL(URL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) GetURLByShortcutURL(shortcutURL string) (*domain.URL, error) {
	resURL, err := us.repo.GetURLByShortcutURL(shortcutURL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) GetURLByOrignalURL(originlURL string) (*domain.URL, error) {
	resURL, err := us.repo.GetURLByShortcutURL(originlURL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) UpdateURL(URL domain.URL) (*domain.URL, error) {
	resURL, err := us.repo.UpdateURL(URL)
	if err != nil {
		return nil, err
	}

	return resURL, nil
}

func (us *URLService) DeleteByShortcutURL(shortcutURL string) error {

	err := us.repo.DeleteByShortcutURL(shortcutURL)
	if err != nil {
		return err
	}

	return nil
}

func (us *URLService) DeleteByOriginalURL(originURL string) error {
	err := us.repo.DeleteByOriginalURL(originURL)
	if err != nil {
		return err
	}

	return nil
}
