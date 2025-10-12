package json

import (
	"encoding/json"

	"github.com/r0kyi/glua/core"
)

type Json struct {
	raw  string
	json map[string]interface{}
}

func (j *Json) encode() error {
	marshal, err := json.Marshal(j.json)
	if err != nil {
		return err
	}
	j.raw = core.B2S(marshal)

	return nil
}

func (j *Json) decode() error {
	err := json.Unmarshal(core.S2B(j.raw), &j.json)
	if err != nil {
		return err
	}

	return nil
}
