package repository

import (
	"context"
	"fmt"
	"gukppap-backend/internal/adapter/repository/mysql/ent"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*ent.Client
}

func NewDB(ctx context.Context) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	clnt, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Migration
	if clnt.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return &DB{clnt}, nil
}
