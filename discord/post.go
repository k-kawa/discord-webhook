package discord

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func Post(webookURL WebhookURL, message *WebhooksPostRequest) error {
	m, err := json.Marshal(message)
	if err != nil {
		return err
	}

	res, err := http.Post(string(webookURL), "application/json", bytes.NewReader(m))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return errors.Errorf("API returned error %v", res)
	}

	return nil
}
