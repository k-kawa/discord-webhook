package we

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

type EnvGcsFile struct {
	Bucket string `json:"bucket"`
	Path   string `json:"path"`
}

func (e *EnvGcsFile) EnvVar() (string, error) {
	//create client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", errors.Errorf("Failed to create Gcs client: %w", err)
	}
	o := client.Bucket(e.Bucket).Object(e.Path)

	r, err := o.NewReader(ctx)
	if err != nil {
		return "", errors.Errorf("Failed to open GCS file gs://%s/%s: %w", e.Bucket, e.Path, err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return "", errors.Errorf("Failed to read GCS file gs://%s/%s: %w", e.Bucket, e.Path, err)
	}
	return string(b), nil
}
