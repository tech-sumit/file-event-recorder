package worker

import (
	"context"
	"fmt"
	"testCode/internal/utils"
	"testCode/internal/utils/taskqueue"
)

func (c *Config) StartWatcher(ctx context.Context) error {
	watcher, err := utils.NewWatcher()
	if err != nil {
		return err
	}
	err = watcher.AddRecursive(c.path)
	if err != nil {
		return err
	}

	for {
		select {
		case event := <-watcher.Events:
			taskqueue.EnqueueTask(func() error {
				err := c.service.RecordFileEvent(ctx, event)
				if err != nil {
					fmt.Printf("WATCHER: Error: %s\n", err.Error())
				}
				return err
			})
			fmt.Println("WATCHER: Added event to queue ", event)
		case <-ctx.Done():
			fmt.Println("WATCHER: Shutting down watcher")
			break
		}
	}
}
