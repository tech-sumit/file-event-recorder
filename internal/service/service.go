package service

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"testCode/internal/repository/sql"
	"time"
)

type Config struct {
	db      *sql.Queries
	NowFunc func() time.Time
}

type Services interface {
	RecordFileEvent(ctx context.Context, event fsnotify.Event) error
}

func New(queries *sql.Queries) *Config {
	return &Config{
		db: queries,
		NowFunc: func() time.Time {
			return time.Now()
		},
	}
}
