package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func (baseConf *baseConfig) parse() error {
	defaultData, err := os.ReadFile("base_config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(defaultData, baseConf)
	if err != nil {
		return err
	}
	return nil
}
