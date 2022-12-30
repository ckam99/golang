package worker

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
	QueueLow      = "low"
)

type TaskProcessor interface {
	// Start new asynq MuxServer
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	db     *sql.DB
}

func NewRedisTaskProcessor(options asynq.RedisClientOpt, db *sql.DB) TaskProcessor {
	server := asynq.NewServer(options, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 10,
			QueueDefault:  5,
			QueueLow:      1,
		},
	})
	return &RedisTaskProcessor{
		server: server,
		db:     db,
	}
}

// Start implements TaskProcessor
func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	if err := processor.server.Start(mux); err != nil {
		return fmt.Errorf("failed to start task processor server: %w", err)
	}
	return nil
}
