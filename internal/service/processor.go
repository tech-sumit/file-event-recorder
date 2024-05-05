package service

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"testCode/internal/repository"
	"testCode/internal/repository/sql"
)

func (c *Config) RecordFileEvent(ctx context.Context, event fsnotify.Event) error {
	var operation string
	switch event.Op {
	case fsnotify.Create:
		operation = repository.CREATED
	case fsnotify.Write:
		operation = repository.UPDATED
	case fsnotify.Rename:
		operation = repository.RENAMED
	case fsnotify.Remove:
		operation = repository.DELETED
	default:
		fmt.Printf("PROCESSOR: Skipped event: Name: %s Event: %s\n", event.Name, event.Op)
		return nil
	}

	var size int64
	if !event.Has(fsnotify.Remove) {
		info, err := os.Stat(event.Name)
		if err != nil {
			panic(err)
		}
		size = info.Size()
	}

	err := c.db.UpsertFile(ctx, sql.UpsertFileParams{
		Path:      event.Name,
		LastEvent: operation,
		Size:      size,
	})

	if err != nil {
		return err
	}

	fmt.Printf("PROCESSOR: Name: %s Event: %s Size: %d\n", event.Name, event.Op, size)
	return nil

}
