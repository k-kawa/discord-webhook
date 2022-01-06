package we

import (
	"reflect"

	"github.com/pkg/errors"
)

type EnvUnion struct {
	Name      string               `json:"name"`
	Raw       *EnvRaw              `json:"raw"`
	GcpSecret *EnvGcpSecretManager `json:"gcp_secret"`
	Gcs       *EnvGcsFile          `json:"gcs"`
}

func (e *EnvUnion) EnvVar() (string, error) {
	var h EnvVarHandler = nil
	for _, v := range []EnvVarHandler{
		e.Raw, e.GcpSecret, e.Gcs,
	} {
		if isNotNil(v) {
			h = v
			break
		}
	}

	if isNil(h) {
		return "", errors.Errorf("No configuration found: %s", e.Name)
	}

	v, err := h.EnvVar()
	if err != nil {
		return "", errors.Errorf("Failed to get value named %s: %w", e.Name, err)
	}
	return v, nil
}

func isNil(v interface{}) bool {
	return (v == nil) || reflect.ValueOf(v).IsNil()
}

func isNotNil(v interface{}) bool {
	return !isNil(v)
}
