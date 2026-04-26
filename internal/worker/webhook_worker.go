package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

// WebhookPayload is the task payload for delivering webhooks.
type WebhookPayload struct {
	URL     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    interface{}       `json:"body"`
}

// WebhookWorker processes webhook tasks.
type WebhookWorker struct {
	logger     *zap.Logger
	httpClient *http.Client
}

// NewWebhookWorker creates a new WebhookWorker.
func NewWebhookWorker(logger *zap.Logger) *WebhookWorker {
	return &WebhookWorker{
		logger: logger,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// ProcessTask handles a webhook:send task.
func (w *WebhookWorker) ProcessTask(_ context.Context, t *asynq.Task) error {
	var payload WebhookPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("unmarshal webhook payload: %w", asynq.SkipRetry)
	}

	body, err := json.Marshal(payload.Body)
	if err != nil {
		return fmt.Errorf("marshal webhook body: %w", asynq.SkipRetry)
	}

	method := payload.Method
	if method == "" {
		method = http.MethodPost
	}

	req, err := http.NewRequest(method, payload.URL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range payload.Headers {
		req.Header.Set(k, v)
	}

	resp, err := w.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("deliver webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("webhook returned status %d", resp.StatusCode)
	}

	w.logger.Info("webhook delivered",
		zap.String("url", payload.URL),
		zap.Int("status", resp.StatusCode),
	)

	return nil
}
