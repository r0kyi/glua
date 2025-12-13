package yaml

import (
	"github.com/goccy/go-yaml"
	"github.com/r0kyi/glua/core"
)

type Yaml struct {
}

func (y *Yaml) encode(yaml_ map[string]any) (string, error) {
	raw, err := yaml.Marshal(yaml_)
	if err != nil {
		return "", err
	}

	return core.B2S(raw), nil
}

func (y *Yaml) decode(raw string) (map[string]any, error) {
	yaml_ := make(map[string]any)
	err := yaml.Unmarshal(core.S2B(raw), &yaml_)
	if err != nil {
		return nil, err
	}

	return yaml_, nil
}
