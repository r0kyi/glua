package ini

import (
	"errors"
	"fmt"

	"gopkg.in/ini.v1"
)

type Ini struct {
}

func (i *Ini) load(filename string) (map[string]any, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, err
	}

	result := make(map[string]any)
	for _, section := range cfg.Sections() {
		sec := make(map[string]any)

		for _, key := range section.Keys() {
			sec[key.Name()] = key.Value()
		}

		result[section.Name()] = sec
	}

	return result, nil
}

func (i *Ini) save(filename string, data map[string]any) error {
	cfg := ini.Empty()
	for name, value := range data {
		section, err := cfg.NewSection(name)
		if err != nil {
			return err
		}

		kv, ok := value.(map[string]any)
		if !ok {
			return errors.New("invalid ini data")
		}

		for k, v := range kv {
			_, err := section.NewKey(k, fmt.Sprint(v))
			if err != nil {
				return err
			}
		}
	}

	err := cfg.SaveTo(filename)
	if err != nil {
		return err
	}

	return nil
}
