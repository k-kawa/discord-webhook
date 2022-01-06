package we

type EnvRaw struct {
	Value string `json:"value"`
}

func (e *EnvRaw) EnvVar() (string, error) {
	return e.Value, nil
}
