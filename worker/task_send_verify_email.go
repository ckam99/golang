package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

// DistributeTaskSendVerifyEmail implements TaskDistributor
func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, p *PayloadSendVerifyEmail, opts ...asynq.Option) error {
	payload, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, payload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}
	log.Println("enqueue task received: ", info)
	return nil
}

// ProcessTaskSendVerifyEmail implements TaskProcessor
func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal task payload: %w", asynq.SkipRetry)
	}
	// TODO check if email exists in DB processor.db.Exec()
	// TODO then send email to user
	log.Println("task processed: ", payload.Username)
	return nil
}
