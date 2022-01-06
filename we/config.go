package we

import (
	"strings"

	"github.com/pkg/errors"
)

type Config interface {
	EnvVars() ([]*EnvVar, error)
	EnvVarStrings() ([]string, error)
}

type configStruct struct {
	envvars []*EnvVar
}

func (c *configStruct) EnvVars() ([]*EnvVar, error) {
	return c.envvars, nil
}

func (c *configStruct) EnvVarStrings() ([]string, error) {
	ret := []string{}

	for _, e := range c.envvars {
		ret = append(ret, strings.Join([]string{e.Name, e.Value}, "="))
	}
	return ret, nil
}

type ConfigFile struct {
	EnvironemntVarilables []*EnvUnion `json:"environment_variables"`
}

func (c *ConfigFile) EnvVars() ([]*EnvVar, error) {
	ret := []*EnvVar{}

	for _, e := range c.EnvironemntVarilables {
		v, err := e.EnvVar()
		if err != nil {
			return nil, errors.Errorf("Failed to get value named %s: %w", e.Name, err)
		}
		ret = append(ret, &EnvVar{
			Name:  e.Name,
			Value: v,
		})
	}

	return ret, nil
}

func (c *ConfigFile) EnvVarStrings() ([]string, error) {
	ret := []string{}

	envvars, err := c.EnvVars()
	if err != nil {
		return nil, err
	}

	for _, e := range envvars {
		ret = append(ret, strings.Join([]string{e.Name, e.Value}, "="))
	}
	return ret, nil
}
