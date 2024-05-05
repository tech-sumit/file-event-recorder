package worker

import (
	"context"
	"testCode/internal/service"
	"time"
)

type Workers interface {
	StartWatcher(context.Context) error
}

type Config struct {
	service *service.Config
	NowFunc func() time.Time
	path    string // Path of directory to watch
}

func New(service *service.Config, path string) *Config {
	return &Config{
		service: service,
		NowFunc: func() time.Time {
			return time.Now()
		},
		path: path,
	}
}
