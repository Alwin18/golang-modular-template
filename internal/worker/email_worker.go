package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

// EmailPayload is the task payload for sending emails.
type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// EmailWorker processes email tasks.
type EmailWorker struct {
	logger *zap.Logger
}

// NewEmailWorker creates a new EmailWorker.
func NewEmailWorker(logger *zap.Logger) *EmailWorker {
	return &EmailWorker{logger: logger}
}

// ProcessTask handles an email:send task.
func (w *EmailWorker) ProcessTask(_ context.Context, t *asynq.Task) error {
	var payload EmailPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("unmarshal email payload: %w", asynq.SkipRetry)
	}

	// TODO: integrate with real email provider (SES, SendGrid, etc.)
	w.logger.Info("sending email",
		zap.String("to", payload.To),
		zap.String("subject", payload.Subject),
	)

	return nil
}
