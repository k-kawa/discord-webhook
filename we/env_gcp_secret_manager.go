package we

import (
	"context"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/pkg/errors"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type EnvGcpSecretManager struct {
	ProjectID string `json:"project_id"`
	Secret    string `json:"secret"`
	Version   string `json:"version"`
}

func (e *EnvGcpSecretManager) EnvVar() (string, error) {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}
	defer client.Close()

	name := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", e.ProjectID, e.Secret, e.Version)
	// Create the request to create the secret.
	accessReq := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	res, err := client.AccessSecretVersion(ctx, accessReq)
	if err != nil {
		return "", errors.Errorf("Failed to access gcp secrets %s: %w", name, err)
	}
	return string(res.Payload.Data), nil
}
