package json

import (
	"encoding/json"

	"github.com/r0kyi/glua/core"
)

type Json struct {
}

func (j *Json) encode(json_ map[string]any) (string, error) {
	marshal, err := json.Marshal(json_)
	if err != nil {
		return "", err
	}
	raw := core.B2S(marshal)

	return raw, nil
}

func (j *Json) decode(raw string) (map[string]any, error) {
	json_ := make(map[string]any)
	err := json.Unmarshal(core.S2B(raw), &json_)
	if err != nil {
		return nil, err
	}

	return json_, nil
}
