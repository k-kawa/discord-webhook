package we

type EnvVar struct {
	Name  string
	Value string
}

type EnvVarHandler interface {
	EnvVar() (string, error)
}
