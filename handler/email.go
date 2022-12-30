package handler

import (
	"database/sql"
	"example-asyncq/worker"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
)

type EmailHandler struct {
	distributor worker.TaskDistributor
	db          *sql.DB
}

func NewEmailHandler(db *sql.DB, distributor worker.TaskDistributor) *EmailHandler {
	return &EmailHandler{
		distributor: distributor,
	}
}

func (h *EmailHandler) SendVerifyEmail(c *fiber.Ctx) error {
	var payload worker.PayloadSendVerifyEmail
	if err := c.ParamsParser(&payload); err != nil {
		return c.Status(422).JSON(fiber.Map{"error": err.Error()})
	}
	opt := []asynq.Option{
		asynq.MaxRetry(3),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	if err := h.distributor.DistributeTaskSendVerifyEmail(c.UserContext(), &payload, opt...); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to distribute task: %s", err),
		})
	}
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("message successfully sent to %s", payload.Username),
	})
}
