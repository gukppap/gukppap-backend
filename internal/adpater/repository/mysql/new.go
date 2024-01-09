package repository

import (
	"context"

	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/guckppap/gukppap-backend/internal/adpater/repository/mysql/ent"
)

type DB struct {
	*ent.Client
}

func NewDB(ctx context.Context) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := drv.DB().Ping(); err != nil {
		return nil, err
	}

	maxPoolIdle, err := strconv.Atoi(os.Getenv("DB_POOL_IDLE_CONN"))
	maxPoolOpen, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_CONN"))
	maxPollLifeTime, err := strconv.Atoi(os.Getenv("DB_POOL_LIFE_TIME"))
	if err != nil {
		return nil, err
	}

	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetMaxOpenConns(maxPoolOpen)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
	clnt := ent.NewClient(ent.Driver(drv))

	// Migration
	if clnt.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return &DB{clnt}, nil
}

func (db *DB) Close() {
	db.Client.Close()
}
