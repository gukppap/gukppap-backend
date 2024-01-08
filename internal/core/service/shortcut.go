package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/guckppap/gukppap-backend/internal/actor/repository/mysql/ent"
	"github.com/guckppap/gukppap-backend/internal/core/domain"
	"github.com/guckppap/gukppap-backend/internal/core/ports"
	"github.com/guckppap/gukppap-backend/pkg/apperrors"
)

// URLService 구조체는 ShortcutService를 구현하는 구현체 입니다.
// 본 구조체는 비지니스 로직을 구현합니다.
// 본 구조체는 ShortcutRepository를 포함합니다.
type UrlService struct {
	shortcutRepository ports.ShortcutRepository
}

func (us *URLService) GetByUrl(ctx context.Context, url string) (*domain.Shortcut, error) {

	shortcut, err := us.shortcutRepository.GetByUrl(ctx, url)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, apperrors.New(apperrors.NotFoundError, err.Error())
		}

		return nil, apperrors.New(apperrors.InternalServerError, err.Error())
	}

	return shortcut, nil
}

func (us *URLService) GetByShortcut(ctx context.Context, shortuct string) (*domain.Shortcut, error) {

	shortcut, err := us.shortcutRepository.GetByShortcut(ctx, shortuct)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, apperrors.New(apperrors.NotFoundError, err.Error())
		}

		return nil, apperrors.New(apperrors.InternalServerError, err.Error())
	}

	return shortcut, nil
}

func (us *URLService) Save(ctx context.Context, url string) error {

	var shortcut string
	var i int
	for i = 1; i <= 5; i++ {
		// uuid를 이용해 shortcut 생성
		shortcut = uuid.NewString()[0:7]

		// GetByShortcut을 성공적으로 수행했다면 (해당 shortcut이 존재한다면) -> err가 nil임.
		if _, err := us.shortcutRepository.GetByShortcut(ctx, shortcut); err != nil {
			break
		}
	}
	// 생성된 5번의 uuid들이 다 중복이었을때는 서버 에러
	if i == 5 {
		//중복된 바로 가기 값이 많아 바로 가기를 만들지 못했습니다.
		return apperrors.New(apperrors.InternalServerError, "Failed to create shortcut due to many duplicate shortcut values.")
	}

	// 새로운 shortcut 생성 (이 shortcut은 domain의 shortcut입니다.)
	newShortcut := domain.NewShortcut(url, shortcut)

	// repostory 의 save 함수 호출 (database에 값 저장)
	if err := us.shortcutRepository.Save(ctx, newShortcut); err != nil {
		return apperrors.New(apperrors.InternalServerError, err.Error())
	}

	// 정상적으로 실행 완료
	return nil
}

func (us *URLService) Update(ctx context.Context, shortcut *domain.Shortcut) error {
	if err := us.shortcutRepository.Update(ctx, shortcut); err != nil {
		return apperrors.New(apperrors.InternalServerError, err.Error())
	}

	return nil
}

func (us *URLService) Delete(ctx context.Context, url string) error {
	if err := us.shortcutRepository.Delete(ctx, url); err != nil {
		return apperrors.New(apperrors.InternalServerError, err.Error())
	}

	return nil
}
