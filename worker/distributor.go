package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(ctx context.Context, p *PayloadSendVerifyEmail, opts ...asynq.Option) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(options asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(options)
	return &RedisTaskDistributor{
		client: client,
	}
}
