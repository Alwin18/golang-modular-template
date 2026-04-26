package worker

import (
	"github.com/hibiken/asynq"
)

// Task type constants.
const (
	TaskSendEmail      = "email:send"
	TaskWebhook        = "webhook:send"
	TaskGenerateReport = "report:generate"
)

// NewEmailTask creates an email task payload.
func NewEmailTask(payload []byte) (*asynq.Task, error) {
	return asynq.NewTask(TaskSendEmail, payload), nil
}

// NewWebhookTask creates a webhook delivery task.
func NewWebhookTask(payload []byte) (*asynq.Task, error) {
	return asynq.NewTask(TaskWebhook, payload), nil
}

// NewReportTask creates a report generation task.
func NewReportTask(payload []byte) (*asynq.Task, error) {
	return asynq.NewTask(TaskGenerateReport, payload), nil
}
