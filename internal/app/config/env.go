package config

import "github.com/caarlos0/env/v6"

func (envConf *envConfig) parse() error {
	err := env.Parse(&envConf)
	if err != nil {
		return err
	}
	return nil
}
