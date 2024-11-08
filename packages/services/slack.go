package services

import (
	"bytes"
	"cma/packages/config"
	"cma/packages/errors"
	"encoding/json"
	"fmt"
	"net/http"
)

type Slack struct {
	webhookURL string
}

func NewSlack() *Slack {
	return &Slack{webhookURL: config.GlobalConfig.SlackWebhookUrl}
}

func (s *Slack) SendMessage(message string) error {
	payload := map[string]string{
		"text": message,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Slack.SendMessage() json marshal error %w", err)
	}

	resp, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("Slack.SendMessage() post request error %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Slack.SendMessage() post request error %w", errors.NewSendMessageStatusCode(resp.StatusCode))
	}

	return nil
}
