package yaml

import (
	"github.com/goccy/go-yaml"
	"github.com/r0kyi/glua/core"
)

type Yaml struct {
	raw  string
	yaml map[string]interface{}
}

func (y *Yaml) encode() error {
	marshal, err := yaml.Marshal(y.yaml)
	if err != nil {
		return err
	}
	y.raw = core.B2S(marshal)

	return nil
}

func (y *Yaml) decode() error {
	err := yaml.Unmarshal(core.S2B(y.raw), &y.yaml)
	if err != nil {
		return err
	}

	return nil
}
