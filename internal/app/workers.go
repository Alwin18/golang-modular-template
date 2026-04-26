package app

import (
	"github.com/Alwin18/golang-module-template/internal/worker"
	"github.com/hibiken/asynq"
)

// startWorkers registers all task handlers and starts the Asynq server.
func (c *Container) startWorkers() {
	mux := asynq.NewServeMux()

	emailWorker := worker.NewEmailWorker(c.Logger)
	webhookWorker := worker.NewWebhookWorker(c.Logger)

	mux.HandleFunc(worker.TaskSendEmail, emailWorker.ProcessTask)
	mux.HandleFunc(worker.TaskWebhook, webhookWorker.ProcessTask)

	if err := c.AsynqServer.Run(mux); err != nil {
		c.Logger.Sugar().Errorf("asynq server error: %v", err)
	}
}
