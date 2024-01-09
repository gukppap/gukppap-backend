package repository

import (
	"context"

	"github.com/guckppap/gukppap-backend/internal/adpater/repository/mysql/ent/shortcut"
	"github.com/guckppap/gukppap-backend/internal/core/domain"
)

type ShortcutRepository struct {
	db *DB
}

func NewShortcutRepository(db *DB) *ShortcutRepository {
	return &ShortcutRepository{db}
}

func (sr *ShortcutRepository) GetByUrl(ctx context.Context, url string) (*domain.Shortcut, error) {
	shortcut, err := sr.db.Shortcut.Query().Where(
		shortcut.URL(url),
	).First(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.Shortcut{
		Url:      shortcut.URL,
		Shortcut: shortcut.Shortcut,
		Exp:      shortcut.Exp,
	}, nil
}

func (sr *ShortcutRepository) GetByShortcut(ctx context.Context, shortuct string) (*domain.Shortcut, error) {
	shortcut, err := sr.db.Shortcut.Query().Where(
		shortcut.Shortcut(shortuct),
	).First(ctx)
	if err != nil {
		return nil, err
	}

	return &domain.Shortcut{
		Url:      shortcut.URL,
		Shortcut: shortcut.Shortcut,
		Exp:      shortcut.Exp,
	}, nil
}

func (sr *ShortcutRepository) Save(ctx context.Context, shortcut *domain.Shortcut) error {
	_, err := sr.db.Shortcut.Create().
		SetURL(shortcut.Url).
		SetShortcut(shortcut.Shortcut).
		SetExp(shortcut.Exp).
		Save(ctx)
	return err
}

func (sr *ShortcutRepository) Update(ctx context.Context, argShortcut, want string) error {
	_, err := sr.db.Shortcut.
		Update().
		Where(shortcut.Shortcut(argShortcut)).
		SetShortcut(want).
		Save(ctx)
	return err
}

func (sr *ShortcutRepository) Delete(ctx context.Context, url string) error {
	_, err := sr.db.Shortcut.Delete().
		Where(shortcut.URL(url)).
		Exec(ctx)

	return err
}
